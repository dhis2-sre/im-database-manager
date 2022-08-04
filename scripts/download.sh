#!/usr/bin/env bash

set -euo pipefail

DATABASE=$1

#$HTTP "$INSTANCE_HOST/databases/$DATABASE/download" "Authorization: Bearer $ACCESS_TOKEN"
curl -H "Authorization: $ACCESS_TOKEN" -L "$INSTANCE_HOST/databases/$DATABASE/download" | cat
