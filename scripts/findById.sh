#!/usr/bin/env bash

set -euo pipefail

source ./auth.sh

DATABASE=$1

$HTTP get "$INSTANCE_HOST/databases/$DATABASE" "Authorization: Bearer $ACCESS_TOKEN"
