#!/usr/bin/env bash

set -euo pipefail

DATABASE=$1
GROUP=$2
NAME=$3

echo "{
  \"name\": \"$NAME\",
  \"groupName\": \"$GROUP\"
}" | $HTTP put "$INSTANCE_HOST/databases/$DATABASE" "Authorization: Bearer $ACCESS_TOKEN"
