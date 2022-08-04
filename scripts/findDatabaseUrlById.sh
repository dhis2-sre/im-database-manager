#!/usr/bin/env bash

set -euo pipefail

DATABASE=$1

$HTTP "$INSTANCE_HOST/databases/$DATABASE/url" "Authorization: Bearer $ACCESS_TOKEN"
