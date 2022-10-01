#!/usr/bin/env bash

set -euo pipefail

source ./auth.sh

$HTTP "$INSTANCE_HOST/databases" "Authorization: Bearer $ACCESS_TOKEN"
