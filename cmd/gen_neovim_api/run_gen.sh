#!/bin/bash

set -e
set -x

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}"  )" && pwd  )"

cd $DIR

go build
./gen_neovim_api -g -o api.go.gen -t api_test.go.gen
cat api.go.gen | gofmt > ../../gen_client_api.go
cat api_test.go.gen | gofmt > ../../gen_client_api_test.go
