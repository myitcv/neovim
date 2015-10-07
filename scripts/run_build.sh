#!/bin/bash

set -e
set -x

eval "$(curl -Ss https://raw.githubusercontent.com/myitcv/bot-ci/temp_nightly/scripts/travis-setup.sh) nightly-x64";

# check that the generated API matches what we have committed
# (this ensures that the Neovim API hasn't moved on without us
# knowing)

# TODO this should be vendored...
go get -u github.com/ncw/gotemplate/...
go get -u github.com/tinylib/msgp
go generate

pushd $TRAVIS_BUILD_DIR/cmd/gen_neovim_api/
go get -d -t -v ./... && go build -v ./...
NEOVIM_BIN=$TRAVIS_BUILD_DIR/_neovim/bin/nvim ./gen_neovim_api -g -o ../../gen_client_api.go -t ../../gen_client_api_test.go
popd

gofmt -w gen_client_api*

output=$(git status --porcelain)

if [ -z "$output"  ]
then
  echo "Git is clean"
else
  git diff
  >&2 echo -e "Git is not clean. The following files should have been committed:\n\n$output"
  exit 1
fi

NEOVIM_BIN=$TRAVIS_BUILD_DIR/_neovim/bin/nvim go test
NEOVIM_BIN=$TRAVIS_BUILD_DIR/_neovim/bin/nvim go test -race
