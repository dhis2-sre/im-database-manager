#!/usr/bin/env bash

set -euo pipefail

ID=$1

$HTTP "$INSTANCE_HOST/databases/$ID/url" "Authorization: Bearer $ACCESS_TOKEN"
