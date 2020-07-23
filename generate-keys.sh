#!/usr/bin/env bash

# Generate CA cert and private key for TLS protocol.
key_dir="$(mktemp -d)"

chmod 0700 "$key_dir"
cd "$key_dir" || exit

openssl req -nodes -new -x509 -keyout ca.key -out ca.crt -subj "/CN=Admission Controller Webhook Demo CA"

openssl genrsa -out tls.key 2048

# Generate a Certificate Signing Request (CSR) for the private key, and sign it with the private key of the CA.
openssl req -new -key tls.key -subj "/CN=webhook-server.webhook-demo.svc" \
    | openssl x509 -req -CA ca.crt -CAkey ca.key -CAcreateserial -out tls.crt


# Create a kubernetes secret from generated .cert and .key files.
kubectl delete secret tls-secret || true
kubectl create secret tls tls-secret \
  --cert "${key_dir}/tls.crt" \
  --key "${key_dir}/tls.key"

# Remove the temp dir.
rm -rf "$key_dir"
