#!/usr/bin/env bash

set -euo pipefail

GROUP=$1
NAME=$2
NEW_NAME=$3
FORMAT=$4

INSTANCE_ID=$($HTTP --check-status "$INSTANCE_HOST/instances-name-to-id/$GROUP/$NAME" "Authorization: Bearer $ACCESS_TOKEN")

echo "{
  \"name\": \"$NEW_NAME\",
  \"format\": \"$FORMAT\"
}" | $HTTP post "$INSTANCE_HOST/databases/save-as/$INSTANCE_ID" "Authorization: Bearer $ACCESS_TOKEN"
