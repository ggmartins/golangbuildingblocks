# MiniBank challenge

#### to build <br>
`docker-compose build`

#### to run <br>
`docker-compose up`

#### to test <br>

cd test; ./test_runset1.sh

```
result.transfer1 (unauthorized)
PASS
PASS
result.login1 (non-existent)
null
PASS
PASS
result.postaccounts (create account) OK
PASS
PASS
result.login1 OK
"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NDg2ODI1NzAsInVzZXJfaWQiOjI3NzYyMzkzMDA1fQ.PJAMeVjwPBIbDVj2FQJF32-WklLyRk8MTacfFuLfDks"
PASS
PASS
result.transfer1 OK
PASS
PASS
result.getaccounts1 OK
PASS
PASS
result.getidbalance OK
PASS
PASS
result.gettranfers1 OK
PASS
PASS
result.getaccounts1 (no token)
PASS
PASS
```

#### TODO
- ways to improve bcrypt with pepper: https://security.stackexchange.com/questions/21263/how-to-apply-a-pepper-correctly-to-bcrypt
- ways to improve jwt with refresh, etc: https://developer.vonage.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr
- postgres, encryption at rest, maybe with https://wiki.postgresql.org/wiki/Transparent_Data_Encryption
- future work: to run on minikube
