See https://github.com/moby/moby/issues/3378

# bergkvist/includeimage

Dockerfile syntax extension for combining multiple Docker images into one.

## Usage

### Syntax

```
INCLUDE <image>:<tag>
```

### Example

```Dockerfile
#syntax=bergkvist/includeimage
FROM alpine:3.12.0
INCLUDE rust:1.44-alpine3.12
INCLUDE python:3.8.3-alpine3.12
```

### How to build

Remember to use Docker buildkit when building, as seen below.

```sh
DOCKER_BUILDKIT=1 docker build -t myimage:latest .
```

## Behavior

- The entire file system of an included image is copied over
- The environment variables of the included image is merged in
  - PATH gets special treatment
- CMD and ENTRYPOINT of included image is ignored.

How it is implemented:

```Dockerfile
# All the file system contents are copied over (using multi-stage builds)
COPY --from=<image>:<tag> / /
# We extract the environment variables from the included image
ENV <-(merge)- docker inspect <image>:<tag> --format='{{.Config.Env}}'
```

## Current Issues/Limitations

- Currently only supports including amd64 linux images.
- The PATH variable will eventually contain a lot of duplicate entries. This could probably be cleaned up, but shouldn't cause any issues. (I think)
