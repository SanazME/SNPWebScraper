#!/usr/bin/env bash
# set -xe
docker image build -t snpedia:1.0 . > build.log # send the stdout to stderr
docker run --rm -v "$PWD":/usr/src/app -w /usr/src/app snpedia:1.0 ./runInDocker.sh
