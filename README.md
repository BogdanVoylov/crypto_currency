# crypto_currency

To create new user or get info about existing one 

link localhost:8080/register

post request

body example(true if user exist false if not)

{
	"Name":"test_user",
	"Registered":true
}

user with name private_key is is already created has a lot of coins(9999....)
user with name tets_user is already created and has 40 coins

To make transaction

link localhost:8080/transfer

post request

body example

{
	"Name":"private_key",
	"Adresses":["test_user"],
	"Amounts":[10]
}

"Name" your name

"Adresses" is an array with name of users that you want to send coins

"Amounts" is an array of amounts, corresponds to array of adresses

