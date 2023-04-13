#!/bin/bash
SECRET_KEY=$(head -c 32 /dev/urandom | base64)

echo "SECRET_KEY=$SECRET_KEY" > .env

echo "Secret Key: $SECRET_KEY"
