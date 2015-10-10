#!/bin/bash

set -e
set -x

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}"  )" && pwd  )"

cd $DIR

api=gen_api.go.gen
api_test=gen_api_test.go.gen

go build
./gen_neovim_api -g -o $api -t $api_test
cat $api | gofmt > ../../gen_client_api.go
cat $api_test | gofmt > ../../gen_client_api_test.go
