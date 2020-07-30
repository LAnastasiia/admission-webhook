#### To run locally:

`./generate-keys-for-local-runs.sh`

`make run ENABLE_WEBHOOKS=true`


#### Tests:
- Unit-test for validating f-ns:  
`go test github.com/LAnastasiia/admission-webhook/pkg/webhooks`

- E2E test  
`cd ./pkg/webhooks/`  
`chmod +x validatingwebhook_test.sh`  
`./validatingwebhook_test.sh`  


#### Metrics:
`curl http://localhost:8080/metrics`