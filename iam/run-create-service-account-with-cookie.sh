#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

go build \
  -o golang-samples \
  -ldflags="
    -X 'main.ServiceAccountName=users/6880568923540230144/serviceAccounts/test-3'
    -X 'main.ServiceAccountDescription=Test'
    -X 'main.ServiceAccountDisplayName=Test'
    -X 'main.Cookie=$(cat .test-cookie)'
    -X 'main.Endpoint=iam.animeapis.dev:443'
  " \
  ./iam/create-service-account-with-cookie

./golang-samples

exit 0