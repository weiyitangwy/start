# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /helloserver

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

CMD ["go", "run", "."]