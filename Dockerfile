FROM golang:1.23.2 AS build

ARG GOOS
ARG PROJECT
ENV GIT_TERMINAL_PROMPT=0
ENV GOARCH=amd64
ENV GOPROXY=https://proxy.golang.org,https://goproxy.io,https://goproxy.dev
ENV GOPRIVATE=github.com/tarmalonchik/*

RUN apt update -y && apt install -y git make g++ bash

WORKDIR /go/src/github.com/service
COPY go.mod ./go.mod
COPY go.sum ./go.sum
COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg
COPY api ./api

RUN go mod download
RUN go test ./...
RUN mkdir bin
RUN go build  -o bin/main cmd/${PROJECT}/main.go

FROM alpine:latest as base
WORKDIR /app
RUN apk add build-base gcompat tcpdump bash nano tzdata iperf3

COPY --from=build /go/src/github.com/service/bin ./bin

CMD ["/app/bin/main"]

