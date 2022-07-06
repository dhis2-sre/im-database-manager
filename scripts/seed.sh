#!/usr/bin/env bash

set -x

GROUP=whoami

function createDatabase() {
  echo "Login..."
  export ACCESS_TOKEN="" && eval $(./login.sh) && echo "$ACCESS_TOKEN"

  echo "Downloading database... $1"
  curl -C - "$2" -o "$HOME/Downloads/$1.sql.gz"
  #curl "$2" -o "$HOME/Downloads/$1.sql.gz"

  echo "Login..."
  export ACCESS_TOKEN="" && eval $(./login.sh) && echo "$ACCESS_TOKEN"

  echo "Uploading database...$1"
  ./upload.sh "$GROUP" "$HOME/Downloads/$1.sql.gz"
}


createDatabase "Sierra Leone - 2.36.0" https://databases.dhis2.org/sierra-leone/2.36.0/dhis2-db-sierra-leone.sql.gz

#createDatabase "Sierra Leone - 2.36.1" https://databases.dhis2.org/sierra-leone/2.36.1/dhis2-db-sierra-leone.sql.gz
#createDatabase "Sierra Leone - 2.36.2" https://databases.dhis2.org/sierra-leone/2.36.2/dhis2-db-sierra-leone.sql.gz
#createDatabase "Sierra Leone - 2.36.3" https://databases.dhis2.org/sierra-leone/2.36.3/dhis2-db-sierra-leone.sql.gz
#createDatabase "Sierra Leone - 2.36.4" https://databases.dhis2.org/sierra-leone/2.36.4/dhis2-db-sierra-leone.sql.gz
#createDatabase "Sierra Leone - 2.36.5" https://databases.dhis2.org/sierra-leone/2.36.5/dhis2-db-sierra-leone.sql.gz
#
#createDatabase "Sierra Leone - 2.36.6" https://databases.dhis2.org/sierra-leone/2.36.6/dhis2-db-sierra-leone.sql.gz
#
#createDatabase "Sierra Leone - 2.37.0" https://databases.dhis2.org/sierra-leone/2.37.0/dhis2-db-sierra-leone.sql.gz
#createDatabase "Sierra Leone - 2.37.1" https://databases.dhis2.org/sierra-leone/2.37.1/dhis2-db-sierra-leone.sql.gz
#createDatabase "Sierra Leone - 2.37.2" https://databases.dhis2.org/sierra-leone/2.37.2/dhis2-db-sierra-leone.sql.gz
#createDatabase "Sierra Leone - 2.37.3" https://databases.dhis2.org/sierra-leone/2.37.3/dhis2-db-sierra-leone.sql.gz
#createDatabase "Sierra Leone - 2.37.4" https://databases.dhis2.org/sierra-leone/2.37.4/dhis2-db-sierra-leone.sql.gz
#createDatabase "Sierra Leone - 2.37.5" https://databases.dhis2.org/sierra-leone/2.37.5/dhis2-db-sierra-leone.sql.gz
#createDatabase "Sierra Leone - 2.37.6" https://databases.dhis2.org/sierra-leone/2.37.6/dhis2-db-sierra-leone.sql.gz

./list.sh | jq '.[].Databases[] | .ID, .Name, .Url'
