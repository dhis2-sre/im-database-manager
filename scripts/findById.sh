#!/usr/bin/env bash

HTTP="http --verify=no --check-status"

ID=$1

$HTTP "$INSTANCE_HOST/databases/$ID" "Authorization: Bearer $ACCESS_TOKEN"
