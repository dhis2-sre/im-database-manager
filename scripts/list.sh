#!/usr/bin/env bash

set -euo pipefail

$HTTP "$INSTANCE_HOST/databases" "Authorization: Bearer $ACCESS_TOKEN"
