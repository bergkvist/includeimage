See https://github.com/moby/moby/issues/3378

# dockerfile-include-syntax

Dockerfile syntax extension for combining multiple Docker images into one.

## Usage

```Dockerfile
#syntax=bergkvist/dockerfile-include-syntax
FROM alpine:3.12.0
INCLUDE rust:1.44-alpine3.12
INCLUDE python:3.8.3-alpine3.12
```

## How does it work?

```Dockerfile
INCLUDE <image>:<tag>
```

How it is implemented:

```Dockerfile
# All the file system contents are copied over (using multi-stage builds)
COPY --from=<image>:<tag> / /

# We also need to include the environment of the image, and merge it into our
# current image. There is no way to do this with standard Dockerfile syntax.
# The PATH variable will get special treatment in this merging process
ENV ???

# We can get the environment of a third-party image by inspecting it:
docker inspect <image>:<tag> --format='{{.Config.Env}}'
```
