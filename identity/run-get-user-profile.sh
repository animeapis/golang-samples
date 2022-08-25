#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

go build \
  -o golang-samples \
  -ldflags="
    -X 'main.User=users/6880568923540230144'
    -X 'main.Endpoint=identity.animeapis.dev:443'
  " \
  ./identity/get-user-profile

./golang-samples

exit 0