#!/usr/bin/env bash

rm public -rf
git add .
git commit -m "update "+`date +%s`
git push
