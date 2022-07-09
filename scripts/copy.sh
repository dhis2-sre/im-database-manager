#!/usr/bin/env bash

set -euo pipefail

NAME=$1
GROUP=$2
SOURCE=$3

echo "{
  \"name\": \"$NAME\",
  \"group\": \"$GROUP\"
}" | $HTTP post "$INSTANCE_HOST/databases/$SOURCE/copy" "Authorization: Bearer $ACCESS_TOKEN"
