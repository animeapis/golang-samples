#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

go build \
  -o golang-samples \
  -ldflags="
    -X 'main.ClientID=$(cat .oauth-client-id)'
    -X 'main.ClientSecret=$(cat .oauth-client-secret)'
    -X 'main.ServiceAccountName=users/6880568923540230144/serviceAccounts/test-4'
    -X 'main.ServiceAccountDescription=Test'
    -X 'main.ServiceAccountDisplayName=Test'
    -X 'main.TokenURL=https://accounts.animeshon.dev/o/oauth2/token'
    -X 'main.Endpoint=iam.animeapis.dev:443'
  " \
  ./iam/create-service-account

./golang-samples

exit 0