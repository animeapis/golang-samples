#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

go build \
  -o golang-samples \
  -ldflags="
    -X 'main.ClientID=$(cat .oauth-client-id)'
    -X 'main.ClientSecret=$(cat .oauth-client-secret)'
    -X 'main.PlaylistParent=users/6880568923540230144'
    -X 'main.PlaylistDisplayName=Test'
    -X 'main.TokenURL=https://accounts.animeshon.dev/o/oauth2/token'
    -X 'main.Endpoint=library.animeapis.dev:443'
  " \
  ./library/create-playlist

./golang-samples

exit 0