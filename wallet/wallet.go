package wallet

import (
	"crypto/rsa"
	"crypto/x509"
	"cryptocurrency/block"
	"cryptocurrency/transaction"
	"encoding/pem"
	"strings"
)

type Wallet struct {
	Adress string
	Key *rsa.PrivateKey
	Account int
	Adresses []string
	Amounts []int
}

func NewWallet(name string, registered bool, chain block.BlockChain) Wallet {

	if(registered){
		publicKey,privateKey := transaction.GetKeys(name)
		parsedKey := PublicKeyToBytes(publicKey)
		strings.Replace(parsedKey, "n", "y",-1)
		var sum int
		for _,c := range chain.Chain{
			for _,y := range c.Data.TxOuts{
				if(y.Adress == parsedKey){
					sum += y.Amount
				}
			}
			for _,z := range c.Data.TxIns{
				if(z.TxOutId == parsedKey){
					for _,x := range c.Data.TxOuts{
						sum -= x.Amount
					}
				}
			}

		}
		//fmt.Println(publicKey)
		buff,_ := transaction.GetKeys("private_key")

		if(string(parsedKey) == string(PublicKeyToBytes(buff))){
			sum = 9999999999999
		}
		return Wallet{Adress: parsedKey, Key:privateKey, Account:sum, Adresses:nil, Amounts:nil}
	}else{
		publicKey,privateKey := transaction.CreateKeys(name)
		return Wallet{Adress: PublicKeyToBytes(publicKey), Key:privateKey, Account:0, Adresses:nil, Amounts:nil}
	}

}
func PublicKeyToBytes(pub *rsa.PublicKey) string {
	pubASN1,_ := x509.MarshalPKIXPublicKey(pub)

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return string(pubBytes)
}

