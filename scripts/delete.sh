#!/usr/bin/env bash

set -euo pipefail

DATABASE=$1

$HTTP delete "$INSTANCE_HOST/databases/$DATABASE" "Authorization: Bearer $ACCESS_TOKEN"
