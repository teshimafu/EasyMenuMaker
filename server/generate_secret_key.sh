#!/bin/bash
SECRET_KEY=$(head -c 32 /dev/urandom | base64)

exit 1 # むやみに更新しないようにexitしている。利用するときはコメントアウトする。
echo "SECRET_KEY=$SECRET_KEY" > .env

echo "Secret Key: $SECRET_KEY"
