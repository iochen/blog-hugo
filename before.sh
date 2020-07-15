#!/usr/bin/env bash

go get -v github.com/iochen/local-search/cmd/localgen/

localgen -p ./content/post/ -o ./content/search/search.json -d "/post/" -l -c