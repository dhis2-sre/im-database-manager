#!/usr/bin/env bash

set -euo pipefail

ID=$1
NAME=$2
GROUP_NAME=$2

echo "{
  \"name\": \"$NAME\",
  \"groupName\": \"$GROUP_NAME\"
}" | $HTTP put "$INSTANCE_HOST/databases/$ID" "Authorization: Bearer $ACCESS_TOKEN"
