#!/usr/bin/env bash

set -euo pipefail

NAME=$1
GROUP_NAME=$2
SOURCE=$3

echo "{
  \"name\": \"$NAME\",
  \"groupName\": \"$GROUP_NAME\"
}" | $HTTP post "$INSTANCE_HOST/databases/$SOURCE/copy" "Authorization: Bearer $ACCESS_TOKEN"
