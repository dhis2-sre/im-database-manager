#!/usr/bin/env bash

HTTP="http --verify=no --check-status"

echo List
$HTTP "$INSTANCE_HOST/databases" "Authorization: Bearer $ACCESS_TOKEN"
