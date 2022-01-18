#!/usr/bin/env bash

set -euo pipefail

ID=$1

$HTTP delete "$INSTANCE_HOST/databases/$ID/unlock" "Authorization: Bearer $ACCESS_TOKEN"
