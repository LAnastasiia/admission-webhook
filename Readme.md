### Create the project

`export PATH=$PATH:/usr/local/kubebuilder/bin`

`kubebuilder init --domain <my.domain>`

`kubebuilder create api --group <admission> --version <v1> --kind <Deployment>`
`kubebuilder create api --group "core" --kind "Pod" --version "v1"`


List of resources & respective Kinds:
https://kubernetes.io/docs/reference/kubectl/overview/#resource-types

`gcr.io/<registry_name>/<image_name>:<tagname>`

To run locally:

`./generate-keys-for-local-runs.sh`

`make run ENABLE_WEBHOOKS=true`

> in new terminal tab 

`curl --request POST -k --header "Content-Type: application/json" --data @sample-request-body.json https://localhost:9443/validate-v1-pod` 