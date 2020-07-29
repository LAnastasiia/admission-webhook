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

// ValidateImageTag checks if containers are defined with proper tags
// e.g. image uri: `gcr.io/<registry_name>/<image_name>:<tagname>`
// For now ValidateImageTag only verifies that the 'latest' tag isn't used so that container won't be using
// an unexpectedly-pushed version of the image when pulling the latest one from image repository.
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
