#!/usr/bin/env bash

set -euo pipefail

ID=$1
FILE=$2

#$HTTP --ignore-stdin --form post "$INSTANCE_HOST/databases/$ID/upload" "database@$FILE" "Authorization: Bearer $ACCESS_TOKEN"
curl --fail --progress-bar -H "Authorization: $ACCESS_TOKEN" -F "database=@$FILE" -L "$INSTANCE_HOST/databases/$ID/upload" | cat
