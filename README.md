# drone-volume-cache

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-volume-cache/status.svg)](http://beta.drone.io/drone-plugins/drone-volume-cache)
[![Join the discussion at https://discourse.drone.io](https://img.shields.io/badge/discourse-forum-orange.svg)](https://discourse.drone.io)
[![Drone questions at https://stackoverflow.com](https://img.shields.io/badge/drone-stackoverflow-orange.svg)](https://stackoverflow.com/questions/tagged/drone.io)
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-volume-cache?status.svg)](http://godoc.org/github.com/drone-plugins/drone-volume-cache)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-volume-cache)](https://goreportcard.com/report/github.com/drone-plugins/drone-volume-cache)
[![](https://images.microbadger.com/badges/image/plugins/volume-cache.svg)](https://microbadger.com/images/plugins/volume-cache "Get your own image badge on microbadger.com")

Drone plugin that allows you to cache directories within the build workspace, this plugin is backed by Docker volumes. For the usage information and a listing of the available options please take a look at [the docs](http://plugins.drone.io/drone-plugins/drone-volume-cache/).

## Build

Build the binary with the following commands:

```
go build
```

## Docker

Build the Docker image with the following commands:

```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-volume-cache
docker build --rm -t plugins/volume-cache .
```

## Usage

Execute from the working directory:

```
docker run --rm \
  -e PLUGIN_FLUSH=true \
  -e PLUGIN_TTL=1 \
  -v $(pwd):$(pwd) \
  -v /tmp/cache:/cache \
  -w $(pwd) \
  plugins/volume-cache

docker run --rm \
  -e PLUGIN_RESTORE=true \
  -e PLUGIN_FILE="backup.tar" \
  -e DRONE_REPO_OWNER="foo" \
  -e DRONE_REPO_NAME="bar" \
  -e DRONE_COMMIT_BRANCH="test"\
  -v $(pwd):$(pwd) \
  -v /tmp/cache:/cache \
  -w $(pwd) \
  plugins/volume-cache

docker run -it --rm \
  -v $(pwd):$(pwd) \
  -v /tmp/cache:/cache \
  -w $(pwd) \
  alpine:latest sh -c "mkdir -p cache && echo 'testing cache' >> cache/test && cat cache/test"

docker run --rm \
  -e PLUGIN_REBUILD=true \
  -e PLUGIN_MOUNT="./cache" \
  -e PLUGIN_FILE="backup.tar" \
  -e DRONE_REPO_OWNER="foo" \
  -e DRONE_REPO_NAME="bar" \
  -e DRONE_COMMIT_BRANCH="test"\
  -v $(pwd):$(pwd) \
  -v /tmp/cache:/cache \
  -w $(pwd) \
  plugins/volume-cache
```
