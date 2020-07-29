#!/usr/bin/env bash

# Generate CA cert and private key for TLS protocol.
key_dir="tls-secret-dir"

rm -rf "$key_dir" || true
mkdir "$key_dir"

chmod 0700 "$key_dir"
cd "$key_dir" || exit

openssl req -nodes -new -x509 -keyout ca.key -out ca.crt -subj "/CN=Admission Controller Webhook Demo CA"

openssl genrsa -out tls.key 2048

# Generate a Certificate Signing Request (CSR) for the private key, and sign it with the private key of the CA.
openssl req -new -key tls.key -subj "/CN=webhook-server.webhook-demo.svc" \
    | openssl x509 -req -CA ca.crt -CAkey ca.key -CAcreateserial -out tls.crt
