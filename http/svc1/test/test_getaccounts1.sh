#!/bin/bash


if [ -f .token ];then
  token=$(cat .token | tr -d '"')
  curl -v --header "Content-Type: application/json" \
    --header "Authorization: Bearer $token" \
    --request GET \
    http://localhost:8000/accounts > result.getaccounts1.json
else
  curl -v --header "Content-Type: application/json" \
    --request GET \
    http://localhost:8000/accounts > result.getaccounts1.json
fi
