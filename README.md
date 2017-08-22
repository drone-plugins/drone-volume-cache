# drone-volume-cache
[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-volume-cache/status.svg)](http://beta.drone.io/api/badges/drone-plugins/drone-volume-cache)

Drone plugin that allows you to `cache` directories within the build workspace. 

## Build

Build the binary with the following commands:

```
go build
```

## Docker

Build the Docker image with the following commands:

```
GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o drone-volume-cache
docker build --rm -t plugins/volume-cache .
```

## Usage

Execute from the working directory:

```
docker run --rm \
  -e PLUGIN_FLUSH=true \
  -v $(pwd):$(pwd) \
  -v /tmp/cache:/cache \
  -w $(pwd) \
  plugins/volume-cache --ttl 1
  
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