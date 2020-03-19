#!/bin/bash

login="$(aws --region=us-east-1 ecr get-login --no-include-email)"
${login}

docker container stop ezgvlcache_1
docker container rm ezgvlcache_1

DOCKERFILEARG="-f docker-compose.yml"

# Darwin = OSX

echo docker-compose $DOCKERFILEARG up -d
docker-compose $DOCKERFILEARG up -d
