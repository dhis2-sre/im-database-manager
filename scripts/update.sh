#!/usr/bin/env bash

HTTP="http --verify=no --check-status"

ID=$1
NAME=$2
GROUP_ID=$3

# Create
echo Create
echo "{
  \"name\": \"$NAME\",
  \"groupId\": $GROUP_ID
}" | $HTTP put "$INSTANCE_HOST/databases/$ID" "Authorization: Bearer $ACCESS_TOKEN"
