#!/bin/bash

set -e

# check that the generated API matches what we have committed
# (this ensures that the Neovim API hasn't moved on without us
# knowing)

pushd $TRAVIS_BUILD_DIR/_cmd/gen_neovim_api/
go get -d -t -v ./... && go build -v ./...
x=`mktemp`
NEOVIM_BIN=$TRAVIS_BUILD_DIR/_neovim/bin/nvim ./gen_neovim_api -g | gofmt > $x
diff $x ../../gen_client_api.go
if [ $? -ne 0 ]
then
  echo "Neovim exposed API differs from committed generated API"
  exit 1
fi
rm $x
popd

NEOVIM_BIN=$TRAVIS_BUILD_DIR/_neovim/bin/nvim go test
NEOVIM_BIN=$TRAVIS_BUILD_DIR/_neovim/bin/nvim go test -race
