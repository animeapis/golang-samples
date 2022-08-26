#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

go build \
  -o golang-samples \
  -ldflags="
    -X 'main.ClientID=$(cat .oauth-client-id)'
    -X 'main.ClientSecret=$(cat .oauth-client-secret)'
    -X 'main.Playlist=users/6880568923540230144/playlists/6962244673655349248'
    -X 'main.Resource1=animes/2576434089677517294/episodes/4211829090873598895'
    -X 'main.Resource2=animes/2576434089677517294/episodes/1109459753516817981'
    -X 'main.Resource3=animes/2576434089677517294/episodes/6916195418543829640'
    -X 'main.TokenURL=https://accounts.animeshon.dev/o/oauth2/token'
    -X 'main.Endpoint=library.animeapis.dev:443'
  " \
  ./library/batch-create-playlist-items

./golang-samples

exit 0