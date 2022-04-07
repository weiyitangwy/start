# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /hello

COPY go.mod ./
COPY hello.go ./

CMD ["go", "run", "."]