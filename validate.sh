#!/usr/bin/env bash

curl -X POST localhost:8080/users/ \
  -H 'Content-Type: application/json' \
  -d '{"email":"foo@bar.baz","password":"my_password"}'

curl localhost:8080/users/ \
  -H 'Content-Type: application/json'

# curl localhost:8080/users/1682794288456 
