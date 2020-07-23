/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:webhook:path=/validate-v1-pod,mutating=false,failurePolicy=fail,groups="",resources=pods,verbs=create;update,versions=v1,name=vpod.kb.io

type imageTagValidator struct {
	Client  client.Client
	decoder *admission.Decoder
}

var restrictedTags = map[string]bool{
	"latest": true,
}

func (v *imageTagValidator) Handle(ctx context.Context, req admission.Request) admission.Response {
	pod := &corev1.Pod{}

	err := v.decoder.Decode(req, pod)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	for _, container := range pod.Spec.Containers {
		containerImageString := container.Image
		i := strings.Index(containerImageString, ":")
		if i > -1 {
			tagString := containerImageString[i+1:]
			if _, keyPresent := restrictedTags[tagString]; keyPresent {
				return admission.Denied(fmt.Sprintf(":%s tag (used in image URI for container '%s') "+
					"is not allowed for security reasons", tagString, container.Name))
			}
		}

	}
	return admission.Allowed("")
}

// imageTagValidator implements admission.DecoderInjector.
// A decoder will be automatically injected.

func (v *imageTagValidator) InjectDecoder(d *admission.Decoder) error {
	v.decoder = d
	return nil
}
