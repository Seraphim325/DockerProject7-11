#!/usr/bin/env bash
set -euo pipefail

SECRETS_DIR="./secrets"

docker secret create db_password ../${SECRETS_DIR}/db_password.txt
docker secret create redis_password ../${SECRETS_DIR}/redis_password.txt