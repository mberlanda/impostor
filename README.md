# Impostor

Impostor is a mock dynamic API server to perform integration testing and accelerate frontend development.

The databases used by the application are in-memory concurrent safe maps.

## Setup

* Install go: https://go.dev/dl/

### Run unit tests

```
go test -v db/*.go
```

### Run with docker

```
docker build -t impostor .
docker run -p 8080:8080 impostor
```

and then visit http://localhost:8080/ping

### Resources

* https://gin-gonic.com/
* https://chemidy.medium.com/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
* https://dev.to/sreramk/go-loadanddelete-and-loadorstore-in-sync-map-why-are-they-needed-30f7
* https://blog.boot.dev/golang/golang-mutex
