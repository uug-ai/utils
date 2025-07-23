FROM golang:1.24.5-bullseye AS builder

ARG github_username
ARG github_token

WORKDIR /build

# Copy the code necessary to build the application
# You may want to change this to copy only what you actually need.
COPY . .

# Get local dependencies from private repo.
RUN git config --global \
    url."https://${github_username}:${github_token}@github.com/".insteadOf \
    "https://github.com/"

# Build the application
RUN go mod download && \
    go build -tags timetzdata,netgo --ldflags '-s -w -extldflags "-static -latomic"' examples/main.go && \
    mkdir -p /utils && \
    mv main /utils 

# Copy or create other directories/files your app needs during runtime.
# E.g. this example uses /data as a working directory that would probably
#      be bound to a perstistent dir when running the container normally
RUN mkdir /data

FROM alpine:latest

COPY --chown=0:0 --from=builder /utils /

RUN apk update && apk add ca-certificates curl libstdc++ libc6-compat --no-cache && rm -rf /var/cache/apk/*

# Set up the app to run as a non-root user inside the /data folder
# User ID 65534 is usually user 'nobody'.
# The executor of this image should still specify a user during setup.
COPY --chown=65534:0 --from=builder /data /data
USER 65534
WORKDIR /data

ENTRYPOINT ["/main"]
