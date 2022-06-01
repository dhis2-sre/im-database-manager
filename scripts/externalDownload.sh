#!/usr/bin/env bash

set -euo pipefail

UUID=$1

curl -L "$INSTANCE_HOST/databases/external/$UUID" | cat
