#!/bin/bash

rm -rf sms-sender
mv ../cmd/sms-sender/sms-sender .
chmod +x sms-sender
docker build --no-cache -t k8s-deploy/sms-sender:v1.2 .
rm -rf sms-sender