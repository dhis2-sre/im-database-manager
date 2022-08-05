#!/usr/bin/env bash

set -euo pipefail

DATABASE=$1
NAME=$2

echo "{
  \"name\": \"$NAME\"
}" | $HTTP put "$INSTANCE_HOST/databases/$DATABASE" "Authorization: Bearer $ACCESS_TOKEN"
