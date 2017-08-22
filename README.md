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

## Configuration Reference

Restore parameters:
* `debug` - Enabled debug. Default false. **optional**.
* `file` -  Tar file that will keep the cache directories. Default /`path`/repo-owner/repo-name/commit-branch.tar. **optional**
* `restore` - If you want to restore this value has to be true. **required**
* `path` - Container path where the volume will be mount. Default /cache. **optional**
* `volumes` - Volume where the tar file will be saved. **required**

Rebuild parameters:
* `debug` - Enabled debug. Default false. **optional**.
* `file` -  Tar file that will keep the cache directories. Default /`path`/repo-owner/repo-name/commit-branch.tar. **optional**
* `mount` - Array of directories to be cached. **required**
* `rebuild` - If you want to rebuild the cache this value has to be true. **required**
* `path` - Container path where the volume will be mount. Default /cache. **optional**
* `volumes` - Volume where the tar file will be saved. **required**

Flush parameters:
* `debug` - Enabled debug. Default false. **optional**.
* `flush` - If you want to flush the cache this value has to be true. **required**
* `ttl` - TTL expressed in days. Purging cached items older then `x` days. 30 days be default. **optional**
* `volumes` - Volume where the tar file will be saved. **required**

## Usage
```
pipeline:

  flush-cache:
    image: plugins/volume-cache
    flush: true
    ttl: 1
    volumes:
      - /var/cache/drone:/cache    

  restore-cache:
    image: plugins/volume-cache
    file: backup.tar
    restore: true
    volumes:
      - /var/cache/drone:/cache
      
  do-something:
    image: alpine:latest
    commands:
      - mkdir -p cache
      - echo $DRONE_COMMIT >> cache/test
      - cat cache/test
      
  rebuild-cache:
    image: plugins/volume-cache
    file: backup.tar
    rebuild: true
    mount:
      - ./cache
    volumes:
      - /var/cache/drone:/cache
```