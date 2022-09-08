#!/usr/bin/env bash

set -euo pipefail

DATABASES=$*

echo "Database(s): $DATABASES"

delete() {
  $HTTP delete "$INSTANCE_HOST/databases/$1" "Authorization: Bearer $ACCESS_TOKEN"
}

for DATABASE in $DATABASES; do
  delete $DATABASE &
done

# shellcheck disable=SC2046
wait $(jobs -p)
