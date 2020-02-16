#!/usr/bin/env bash

set -x 

docker run --rm -v "$PWD":/usr/src/app -w /usr/src/app golang ./runInDocker.sh
# ./runInDocker.sh
