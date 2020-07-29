### To create the project

`export PATH=$PATH:/usr/local/kubebuilder/bin`

`kubebuilder init --domain <my.domain>`

`kubebuilder create api --group <admission> --version <v1> --kind <Deployment>`
`kubebuilder create api --group "core" --kind "Pod" --version "v1"`


#### To run locally:

`./generate-keys-for-local-runs.sh`

`make run ENABLE_WEBHOOKS=true`

In new terminal tab: 

`cd ./pkg/webhooks/testdata`
`curl --request POST -k --header "Content-Type: application/json" --data @admission-review.json https://localhost:9443/validate-v1-pod` 


#### Tests:
- Unit-test for validating f-ns:  
`go test github.com/LAnastasiia/admission-webhook/pkg/webhooks`

- E2E test  
`cd ./pkg/webhooks/`  
`chmod +x validatingwebhook_test.sh`  
`./validatingwebhook_test.sh`  