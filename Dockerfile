# syntax = docker/dockerfile:experimental
FROM golang:1.14 AS build

ARG IGNORECACHE=0

ADD . /tmpdns
RUN --mount=type=cache,target=/go \
    cd /tmpdns \
    && echo "go get" \
    && go get -d \
    && echo "go build" \
    && GOCACHE=/go/.cache CGO_ENABLED=0 go build -ldflags='-s -w'

FROM alpine:3.10
EXPOSE 53/udp
COPY --from=build /tmpdns/tmpdns /
ENTRYPOINT ["/tmpdns"]

