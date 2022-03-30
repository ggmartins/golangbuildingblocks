#!/bin/bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"cpf":"27762393005","secret":"minibank1"}' \
  http://localhost:8000/login | jq '.token' | tee .token

#"create_at":"2022-03-15 15:19:03-04"}
