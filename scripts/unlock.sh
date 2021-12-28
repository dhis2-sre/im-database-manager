#!/usr/bin/env bash

HTTP="http --verify=no --check-status"

ID=$1

echo Unlock database
$HTTP delete "$INSTANCE_HOST/databases/$ID/unlock" "Authorization: Bearer $ACCESS_TOKEN"
