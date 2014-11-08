#!/bin/bash

set -e

# Use the old approach to building until https://github.com/neovim/neovim/pull/1335 is merged
source <(curl -fsSL https://raw.githubusercontent.com/neovim/neovim/master/.ci/common.sh)

set_environment /opt/neovim-deps

if [ ! -d $TRAVIS_BUILD_DIR/_neovim ]
then
  git clone -b provider-autocmd https://github.com/tarruda/neovim.git $TRAVIS_BUILD_DIR/_neovim
fi

pushd $TRAVIS_BUILD_DIR/_neovim
# git fetch origin
# git rebase origin/master
make
popd

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
