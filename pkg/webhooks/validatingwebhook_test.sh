#!/usr/bin/env bash


Test () {
  testRequestFile="$1"
  expectedReturn="$2"

  receivedResponse="$(curl --request POST -k --header "Content-Type: application/json" --data "@$testRequestFile" https://localhost:9443/validate-v1-pod)"
  receivedReturn=$(echo "$receivedResponse" | jq '. | .response.allowed')

  if [[ "$expectedReturn" == "$receivedReturn" ]]; then
    echo "OK."
  else
    echo "!! TEST FAILED"
    exit 1
  fi
}


Test "./testdata/admission-review-with-correct-image-tags.json" "true"
Test "./testdata/admission-review-with-badly-tagged-container-images.json" "false"
exit 0