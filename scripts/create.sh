#!/usr/bin/env bash

set -euo pipefail

NAME=$1
GROUP_NAME=$2

echo "{
  \"name\": \"$NAME\",
  \"groupName\": \"$GROUP_NAME\"
}" | $HTTP post "$INSTANCE_HOST/databases" "Authorization: Bearer $ACCESS_TOKEN"
