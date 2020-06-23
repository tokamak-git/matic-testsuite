#!/usr/bin/env sh

CWD=$PWD

for i in {0..4}
do
  bash ./docker-heimdall-start.sh $i
done
