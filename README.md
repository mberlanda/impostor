# Impostor

Impostor is a mock dynamic API server to perform integration testing and accelerate frontend development.

## Setup

* Install go: https://go.dev/dl/

### Run with docker

```
docker build -t impostor .
docker run -p 8080:8080 impostor
```

and then visit http://localhost:8080/ping

### Resources

* https://gin-gonic.com/
* https://chemidy.medium.com/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
