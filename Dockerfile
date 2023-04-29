############################
# STEP 1 build executable binary
############################
FROM golang:1.20-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/github.com/mberlanda/impostor
COPY . .

WORKDIR $GOPATH/src/github.com/mberlanda/impostor/api
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/impostor-api

############################
# STEP 2 build a small image
############################
FROM scratch
COPY --from=builder /go/bin/impostor-api /go/bin/impostor-api

EXPOSE 8080
ENTRYPOINT ["/go/bin/impostor-api"]