# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /helloserver

COPY go.mod ./
COPY helloserver.go ./

CMD ["go", "run", "."]