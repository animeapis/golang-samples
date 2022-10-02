#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

go build \
  -o golang-samples \
  -ldflags="
    -X 'main.ClientID=$(cat .oauth-client-id)'
    -X 'main.ClientSecret=$(cat .oauth-client-secret)'
    -X 'main.ContributionName=users/6880568923540230144/contributions/ca81a745-67e1-4107-a296-f4e550a99bfe'
    -X 'main.TokenURL=https://accounts.animeshon.dev/o/oauth2/token'
    -X 'main.Endpoint=knowledge.animeapis.dev:443'
  " \
  ./knowledge/get-contribution

./golang-samples

exit 0