#!/usr/bin/env bash

set -euo pipefail

DATABASE=$1

$HTTP delete "$INSTANCE_HOST/databases/$DATABASE/unlock" "Authorization: Bearer $ACCESS_TOKEN"
