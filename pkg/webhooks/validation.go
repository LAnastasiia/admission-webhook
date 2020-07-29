package webhooks

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
	"strings"
)

var restrictedTags = map[string]bool{
	"latest": true,
}

func ValidateImageTag(pod corev1.Pod) admission.Response {
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
