#!/bin/bash

set -e

if [ ! -d $TRAVIS_BUILD_DIR/_neovim ]
then
  git clone https://github.com/neovim/neovim.git $TRAVIS_BUILD_DIR/_neovim
fi

pushd $TRAVIS_BUILD_DIR/_neovim
git fetch origin
git rebase origin/master
make
popd

go get -d -v ./... && go build -v ./...

NEOVIM_BIN=$TRAVIS_BUILD_DIR/_neovim/build/bin/nvim go test
