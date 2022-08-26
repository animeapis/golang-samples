#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

go build \
  -o golang-samples \
  -ldflags="
    -X 'main.ClientID=$(cat .oauth-client-id)'
    -X 'main.ClientSecret=$(cat .oauth-client-secret)'
    -X 'main.Code=$(cat .test-oauth2-code)'
    -X 'main.State=$(cat .test-oauth2-state)'
    -X 'main.Flow=users/6880568923540230144/flows/myanimelist-net'
    -X 'main.TokenURL=https://accounts.animeshon.dev/o/oauth2/token'
    -X 'main.Endpoint=credentials.animeapis.dev:443'
  " \
  ./credentials/exchange

./golang-samples

exit 0