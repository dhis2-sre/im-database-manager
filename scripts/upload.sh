#!/usr/bin/env bash

set -euo pipefail

GROUP=$1
FILE=$2

#$HTTP --ignore-stdin --form post "$INSTANCE_HOST/databases/$ID/upload" "group=$GROUP" "database@$FILE" "Authorization: Bearer $ACCESS_TOKEN"
curl --fail --progress-bar -H "Authorization: $ACCESS_TOKEN" -F "group=$GROUP" -F "database=@$FILE" -L "$INSTANCE_HOST/databases" | cat
