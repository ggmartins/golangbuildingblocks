curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"id":9001,"account_origin_id":1001,"account_destination_id":1002,"amount":4.50,"create_at":"2022-03-15 15:19:03-04"}' \
  http://localhost:8000/transfers
