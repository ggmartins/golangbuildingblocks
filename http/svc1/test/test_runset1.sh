#!/bin/bash


###
echo "result.transfer1 (unauthorized)"

#add expired token to test Unauthorized
echo "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NDg2Mzg4NDUsInVzZXJfaWQiOjI3NzYyMzkzMDA1fQ.2G_Lh4PB4GNVCunZbcNtBBkWd5HhTF0bOWDseV9Rpw0" > .token

./test_transfer1.sh 2>result.transfer1.txt

if cat result.transfer1.json | grep -q "Desautorizado"; then  echo PASS; else echo FAIL; fi
if cat result.transfer1.txt | grep -q "HTTP/1.1 401 Unauthorized"; then  echo PASS; else echo FAIL; fi

###
echo "result.login1 (non-existent)"

./test_login1.sh 2>result.login1.txt

if cat result.login1.json | grep -q "Negado"; then  echo PASS; else echo FAIL; fi
if cat result.login1.txt | grep -q "HTTP/1.1 401 Unauthorized"; then  echo PASS; else echo FAIL; fi


###
echo "result.postaccounts (create account) OK"

./test_createaccount1.sh 2>result.createaccount1.txt

if cat result.createaccount1.json | grep -q "Aprovado"; then  echo PASS; else echo FAIL; fi
if cat result.createaccount1.txt | grep -q "HTTP/1.1 200 OK"; then  echo PASS; else echo FAIL; fi

###
echo "result.login1 OK"

./test_login1.sh 2>result.login1.txt

if cat result.login1.json | grep -q "Aprovado"; then  echo PASS; else echo FAIL; fi
if cat result.login1.txt | grep -q "HTTP/1.1 200 OK"; then  echo PASS; else echo FAIL; fi


###
echo "result.transfer1 OK"

./test_transfer1.sh 2>result.transfer1.txt

if cat result.transfer1.json | grep -q "Aprovado"; then  echo PASS; else echo FAIL; fi
if cat result.transfer1.txt | grep -q "HTTP/1.1 200 OK"; then  echo PASS; else echo FAIL; fi

###
echo "result.getaccounts1 OK"

./test_getaccounts1.sh 2>result.getaccounts1.txt

if cat result.getaccounts1.json | grep -q "Samantha"; then  echo PASS; else echo FAIL; fi
if cat result.getaccounts1.txt | grep -q "HTTP/1.1 200 OK"; then  echo PASS; else echo FAIL; fi

###
echo "result.gettranfers1 OK"

./test_gettransfers1.sh 2>result.gettransfers1.txt

if cat result.gettransfers1.json | grep -q "4.5"; then  echo PASS; else echo FAIL; fi
if cat result.gettransfers1.txt | grep -q "HTTP/1.1 200 OK"; then  echo PASS; else echo FAIL; fi

###
echo "result.getaccounts1 (no token)"

rm .token
./test_getaccounts1.sh 2>result.getaccounts1.txt

if cat result.getaccounts1.json | grep -q "Negado"; then  echo PASS; else echo FAIL; fi
if cat result.getaccounts1.txt | grep -q "HTTP/1.1 401 Unauthorized"; then  echo PASS; else echo FAIL; fi


