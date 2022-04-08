# start
Get your Go and Docker installed.

build with below command

```
docker build -t helloserver .
```

run with below command

```
docker run -d -p 8080:8080 helloserver
```

To run test

```
go test -v ./...
```

WIP: this test program although starts a container, but it was not listening to port 8080, hence the container needs to be running for the later part of the test to request from.
