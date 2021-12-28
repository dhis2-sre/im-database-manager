#!/usr/bin/env bash

HTTP="http --verify=no --check-status"

ID=$1
FILE=$2

$HTTP --ignore-stdin --form post "$INSTANCE_HOST/databases/$ID/upload" "database@$FILE" "Authorization: Bearer $ACCESS_TOKEN"
