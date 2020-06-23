#!/usr/bin/env sh

CWD=$PWD

for i in {0..4}
do
  bash ./docker-bor-start.sh $i
done
