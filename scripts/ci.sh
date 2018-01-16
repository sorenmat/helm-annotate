#!/bin/bash -ex

WORKSPACE=/go/src/github.com/Tradeshift/$1

docker run \
	--volume $(pwd):$WORKSPACE \
	--workdir $WORKSPACE \
	--rm \
golang:1.9 scripts/test.sh $1
#docker build -t docker.tradeshift.net/$1:${BRANCH_NAME}-${BUILD_NUMBER} .
#docker push docker.tradeshift.net/$1:${BRANCH_NAME}-${BUILD_NUMBER}
