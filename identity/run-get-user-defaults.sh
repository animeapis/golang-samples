#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

go build \
  -o golang-samples \
  -ldflags="
    -X 'main.ClientID=$(cat .oauth-client-id)'
    -X 'main.ClientSecret=$(cat .oauth-client-secret)'
    -X 'main.User=users/6880568923540230144'
    -X 'main.TokenURL=https://accounts.animeshon.dev/o/oauth2/token'
    -X 'main.Endpoint=identity.animeapis.dev:443'
  " \
  ./identity/get-user-defaults

./golang-samples

exit 0