#!/usr/bin/env bash

set -euo pipefail

ID=$1

$HTTP "$INSTANCE_HOST/databases/$ID" "Authorization: Bearer $ACCESS_TOKEN"
