#!/usr/bin/env bash

HTTP="http --verify=no --check-status"

NAME=$1
GROUP_ID=$2

echo "{
  \"name\": \"$NAME\",
  \"groupId\": $GROUP_ID
}" | $HTTP post "$INSTANCE_HOST/databases" "Authorization: Bearer $ACCESS_TOKEN"
