#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

go build \
  -o golang-samples \
  -ldflags="
    -X 'main.ClientID=$(cat .oauth-client-id)'
    -X 'main.ClientSecret=$(cat .oauth-client-secret)'
    -X 'main.Tracker=users/6880568923540230144/trackers/6962274256758145024'
    -X 'main.Resource=animes/2576434089677517294/episodes/4211829090873598895'
    -X 'main.TokenURL=https://accounts.animeshon.dev/o/oauth2/token'
    -X 'main.Endpoint=tracker.animeapis.dev:443'
  " \
  ./tracker/create-activity

./golang-samples

exit 0