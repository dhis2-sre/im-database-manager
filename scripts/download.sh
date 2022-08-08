#!/usr/bin/env bash

set -euo pipefail

DATABASE=$1

curl -H "Authorization: $ACCESS_TOKEN" -L "$INSTANCE_HOST/databases/$DATABASE/download" | cat
