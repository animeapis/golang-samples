#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

go build \
  -o golang-samples \
  -ldflags="
    -X 'main.ClientID=$(cat .oauth-client-id)'
    -X 'main.ClientSecret=$(cat .oauth-client-secret)'
    -X 'main.ImageAnnotationParent=users/6880568923540230144'
    -X 'main.Resource=//image.animeapis.com/users/6880568923540230144/albums/6956454702873989120/images/6956455742973300736'
    -X 'main.TokenURL=https://accounts.animeshon.dev/o/oauth2/token'
    -X 'main.Endpoint=vision.animeapis.dev:443'
  " \
  ./vision/create-image-annotation

./golang-samples

exit 0