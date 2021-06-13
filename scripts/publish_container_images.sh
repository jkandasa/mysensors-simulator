#!/bin/bash

source ./scripts/version.sh

# container registry
REGISTRY='docker.io/jkandasa'
IMAGE_NAME="mysensors-simulator"
PLATFORMS="linux/arm/v6,linux/arm/v7,linux/arm64,linux/amd64,linux/386"
IMAGE_TAG=${VERSION}

# debug lines
echo $PWD
ls -alh
git branch

# build and push to quay.io
docker buildx build --push \
  --progress=plain \
  --build-arg=GOPROXY=${GOPROXY} \
  --platform ${PLATFORMS} \
  --file Dockerfile \
  --tag ${REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG} .


# build and push to docker hub
docker buildx build --push \
  --progress=plain \
  --build-arg=GOPROXY=${GOPROXY} \
  --platform ${PLATFORMS} \
  --file Dockerfile \
  --tag ${ALT_REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG} .
