module github.com/LAnastasiia/admission-webhook

go 1.13

require (
	github.com/go-logr/logr v0.1.0
	github.com/go-logr/zapr v0.1.0
	go.uber.org/zap v1.10.0
	k8s.io/api v0.17.2
	k8s.io/apiextensions-apiserver v0.17.2
	k8s.io/apimachinery v0.17.2
	k8s.io/client-go v0.17.2
	sigs.k8s.io/controller-runtime v0.5.0
)
