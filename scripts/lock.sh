#!/usr/bin/env bash

HTTP="http --verify=no --check-status"

ID=$1
INSTANCE_ID=$2

echo Lock database
echo "{
  \"instanceId\": $INSTANCE_ID
}" | $HTTP "$INSTANCE_HOST/databases/$ID/lock" "Authorization: Bearer $ACCESS_TOKEN"
