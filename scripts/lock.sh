#!/usr/bin/env bash

set -euo pipefail

DATABASE_ID=$1
INSTANCE_ID=$2

echo "{
  \"instanceId\": $INSTANCE_ID
}" | $HTTP "$INSTANCE_HOST/databases/$DATABASE_ID/lock" "Authorization: Bearer $ACCESS_TOKEN"
