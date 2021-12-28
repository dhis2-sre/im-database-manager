#!/usr/bin/env bash

HTTP="http --verify=no --check-status"

ID=$1

echo Find by id
$HTTP "$INSTANCE_HOST/databases/$ID" "Authorization: Bearer $ACCESS_TOKEN"
