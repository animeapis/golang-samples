#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

go build \
  -o golang-samples \
  -ldflags="
    -X 'main.ClientID=$(cat .oauth-client-id)'
    -X 'main.ClientSecret=$(cat .oauth-client-secret)'
    -X 'main.Album=users/6880568923540230144/albums/6956454702873989120'
    -X 'main.ImageURL=https://avatars.githubusercontent.com/u/48774482'
    -X 'main.TokenURL=https://accounts.animeshon.dev/o/oauth2/token'
    -X 'main.Endpoint=image.animeapis.dev:443'
  " \
  ./image/import-image

./golang-samples

exit 0