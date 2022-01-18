#!/usr/bin/env bash

set -euo pipefail

ID=$1
NAME=$2
GROUP_NAME=$2

GROUP_ID=$($HTTP "$INSTANCE_HOST/groups-name-to-id/$GROUP_NAME" "Authorization: Bearer $ACCESS_TOKEN")

echo "{
  \"name\": \"$NAME\",
  \"groupId\": $GROUP_ID
}" | $HTTP put "$INSTANCE_HOST/databases/$ID" "Authorization: Bearer $ACCESS_TOKEN"
