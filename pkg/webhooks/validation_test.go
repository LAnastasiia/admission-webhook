package webhooks_test

import (
	"github.com/LAnastasiia/admission-webhook/pkg/webhooks"
	"k8s.io/api/admission/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"reflect"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
	"testing"
)

var testData = []struct {
	examplePods    corev1.Pod
	expectedReturn admission.Response
}{
	// Image tags used in all of the pod's containers are allowed.
	{
		corev1.Pod{
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{Image: "some-repo-domain/image-name-1:image-tag"},
					{Image: "some-repo-domain/image-name-2:image-tag"},
				},
			},
		},
		admission.Response{
			AdmissionResponse: v1beta1.AdmissionResponse{Allowed: true},
		},
	},
	// One of containers contains a restricted image tag.
	{
		corev1.Pod{
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{Image: "some-repo-domain/image-name-1:image-tag"},
					{Image: "some-repo-domain/image-name-2:latest"},
				},
			},
		},
		admission.Response{
			AdmissionResponse: v1beta1.AdmissionResponse{Allowed: false},
		},
	},
	// Pod does not have any containers.
	{
		corev1.Pod{
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{},
			},
		},
		admission.Response{
			AdmissionResponse: v1beta1.AdmissionResponse{Allowed: true},
		},
	},
	// All of the pod's containers have restricted image tags.
	{
		corev1.Pod{
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{Image: "some-repo-domain/image-name-1:latest"},
					{Image: "some-repo-domain/image-name-2:latest"},
				},
			},
		},
		admission.Response{
			AdmissionResponse: v1beta1.AdmissionResponse{Allowed: false},
		},
	},
}

// TestImageTagValidation checks the logic of ValidateImageTag function container tags admission.
func TestImageTagValidation(t *testing.T) {
	for _, td := range testData {
		actualReturn := webhooks.ValidateImageTag(td.examplePods)
		if reflect.DeepEqual(actualReturn, td.expectedReturn) {
			t.Errorf("got %t, want %t", actualReturn.Allowed, td.expectedReturn.Allowed)
		}
	}
}
