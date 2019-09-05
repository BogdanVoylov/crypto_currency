package transaction

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"strconv"
)

type UnspentTxOut struct {
	TxOutId    string
	TxOutIndex int
	Address    string
	Amount     int
}
type TxIn struct {
	TxOutId    string
	TxOutIndex int
	Signature  []byte
}

type TxOut struct {
	Adress string
	Amount int
}

type Transaction struct {
	Id []byte
	TxIns []TxIn
	TxOuts []TxOut
}

func NewTransaction(key *rsa.PrivateKey, data UnspentTxOut, adresses []string, amounts []int) Transaction{
	var txIn []TxIn
	var txOut []TxOut
	for i,u := range adresses{
		txIn = append(txIn,TxIn{data.Address,i, nil})
		txOut = append(txOut, TxOut{u,amounts[i]})
	}
	transactionId := GetTransactionId(txIn,txOut)
	for _,y := range txIn{
		y.Signature = CreateSignature(key, transactionId)
	}
	return Transaction{Id:transactionId, TxIns:txIn, TxOuts:txOut}
}

func GetTransactionId(txIns []TxIn, txOuts []TxOut) []byte{
	a := ""
	for _,i := range txIns{
		a += strconv.Itoa(i.TxOutIndex)
		a += i.TxOutId
	}
	b := ""
	for _,i := range txOuts{
		b += strconv.Itoa(i.Amount)
		b += i.Adress
	}
	buf := sha256.Sum256([]byte(a+b))
	return buf[0:31]
}

func CreateSignature(key *rsa.PrivateKey, id []byte) []byte{
	rng := rand.Reader
	dataToSign := id
	signature,_ := rsa.SignPKCS1v15(rng, key, crypto.SHA256, dataToSign)
	return signature
}