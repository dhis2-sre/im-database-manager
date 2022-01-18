#!/usr/bin/env bash

set -euo pipefail

ID=$1
INSTANCE_ID=$2

echo "{
  \"instanceId\": $INSTANCE_ID
}" | $HTTP "$INSTANCE_HOST/databases/$ID/lock" "Authorization: Bearer $ACCESS_TOKEN"
