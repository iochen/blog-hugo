#!/usr/bin/env bash

# for local search
go get -v github.com/iochen/local-search/cmd/localgen/
localgen -p ./content/post/ -o ./content/search/search.json -d "/post/" -l -c

# for cdn
sed -i '/^#.*FOR_CDN/s/^#//' config.toml