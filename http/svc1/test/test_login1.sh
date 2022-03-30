#!/bin/bash
curl -v --header "Content-Type: application/json" \
  --request POST \
  --data '{"cpf":"27762393005","secret":"minibank1"}' \
  http://localhost:8000/login > result.login1.json

cat result.login1.json | jq '.token' | tee .token

#"create_at":"2022-03-15 15:19:03-04"}
