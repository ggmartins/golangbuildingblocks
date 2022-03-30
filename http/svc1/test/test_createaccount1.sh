#!/bin/bash

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"id":3001,"name":"Tylor Hawkins","cpf":"27762393005","secret":"minibank1","balance":0.0}' \
  http://localhost:8000/accounts

#"create_at":"2022-03-15 15:19:03-04"}
