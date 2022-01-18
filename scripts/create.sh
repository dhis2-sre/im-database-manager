#!/usr/bin/env bash

set -euo pipefail

NAME=$1
GROUP_NAME=$2

GROUP_ID=$($HTTP "$INSTANCE_HOST/groups-name-to-id/$GROUP_NAME" "Authorization: Bearer $ACCESS_TOKEN")

echo "{
  \"name\": \"$NAME\",
  \"groupId\": $GROUP_ID
}" | $HTTP post "$INSTANCE_HOST/databases" "Authorization: Bearer $ACCESS_TOKEN"
