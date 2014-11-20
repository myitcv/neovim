#!/bin/bash

set -e
set -x

eval "$(curl -Ss https://raw.githubusercontent.com/neovim/bot-ci/master/scripts/travis-setup.sh) nightly-x64";

# check that the generated API matches what we have committed
# (this ensures that the Neovim API hasn't moved on without us
# knowing)

pushd $TRAVIS_BUILD_DIR/_cmd/gen_neovim_api/
go get -d -t -v ./... && go build -v ./...
x=`mktemp`
NEOVIM_BIN=$TRAVIS_BUILD_DIR/_neovim/bin/nvim ./gen_neovim_api -g -o api.go.gen -t api_test.go.gen
cat api.go.gen | gofmt | diff -- - ../../gen_client_api.go
if [ $? -ne 0 ]
then
  echo "Neovim exposed API differs from committed generated API"
  exit 1
fi
cat api_test.go.gen | gofmt | diff -- - ../../gen_client_api_test.go
if [ $? -ne 0 ]
then
  echo "Neovim exposed API differs that resulted in different test interfaces"
  exit 1
fi
rm $x
popd

NEOVIM_BIN=nvim go test
NEOVIM_BIN=nvim go test -race
