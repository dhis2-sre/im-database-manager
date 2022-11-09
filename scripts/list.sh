#!/usr/bin/env bash

set -euo pipefail

source ./auth.sh

$HTTP get "$INSTANCE_HOST/databases" "Authorization: Bearer $ACCESS_TOKEN"
