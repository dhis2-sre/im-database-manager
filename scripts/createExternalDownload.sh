#!/usr/bin/env bash

set -euo pipefail

DATABASE_ID=$1
EXPIRATION=$2

echo "{
  \"expiration\": \"$EXPIRATION\"
}" | $HTTP post "$INSTANCE_HOST/databases/$DATABASE_ID/external" "Authorization: Bearer $ACCESS_TOKEN"
