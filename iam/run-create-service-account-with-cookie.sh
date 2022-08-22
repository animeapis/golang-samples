#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

go build \
  -o golang-samples \
  -ldflags="
    -X 'main.ServiceAccountName=users/6880568923540230144/serviceAccounts/test-3'
    -X 'main.ServiceAccountDescription=Test'
    -X 'main.ServiceAccountDisplayName=Test'
    -X 'main.Cookie=ory_kratos_session=MTY2MTEyMDc5OHw0Qkx1NC1BU09MOXd1OVdrcTNuR052blhRVEZwN19rSzJpNzZQcVllcVV6anUyTUZmTk9aUjd4NmZpOEFxckhabWZ4WTNzQXgwTXNJUTBLYlFZc190ZndienBBTmp0SUxiMWwzUlBRRW9KRTBTN2E3SnZiQ3Ficlc2eWVjeDg4akNNVlFUSHAtUFE9PXzCXTQ6LS3EBB9YXc1IZ_jtbsERBK7_wcSB8LDlNXHjKQ=='
    -X 'main.Endpoint=iam.animeapis.dev:443'
  " \
  ./iam/create-service-account-with-cookie

./golang-samples

exit 0