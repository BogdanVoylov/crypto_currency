package main

import (
	"crypto/rsa"
	"crypto/sha256"
	"cryptocurrency/block"
	"cryptocurrency/transaction"
	"cryptocurrency/wallet"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type registerRequest struct {
	Name string
	Registered bool
}

type transferRequest struct{
	Name string
	Adresses []string
	Amounts []int
}

var BlockChain block.BlockChain
var Account wallet.Wallet

func main() {
	byteValue, _ := ioutil.ReadFile("test.json")
	json.Unmarshal(byteValue, &BlockChain)
	fmt.Println(string(BlockChain.GetJSONBlockChain()))
	i := []int{1}
	fmt.Println(len(i))
	fmt.Println(len(BlockChain.Chain))
	http.HandleFunc("/register", balance)
	http.HandleFunc("/transfer", transfer)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	/*blockChain := block.NewBlockChain()

	wallet := wallet2.NewWallet("private_key", true)
	fmt.Println(string(wallet.Adress))*/
}

func balance(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var t registerRequest
	err := decoder.Decode(&t)
	cli,_ := json.Marshal(t)
	fmt.Println(string(cli))
	Account = wallet.NewWallet(t.Name, t.Registered, BlockChain)
	if err != nil {
		panic(err)
	}
	//fmt.Fprintf(w, string(Account.Account))
	res,_ := json.Marshal(Account)
	fmt.Fprintf(w, string(res))
}

func transfer(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var t transferRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	Account = wallet.NewWallet(t.Name, true, BlockChain)
	check := 0
	for _,i :=range t.Amounts{
		check+=i
	}
	if(check < Account.Account || check == Account.Account){
		var adress *rsa.PublicKey
		for _,a := range t.Adresses{
			adress,_ = transaction.GetKeys(a)
			Account.Adresses = append(Account.Adresses,wallet.PublicKeyToBytes(adress))
		}

		Account.Amounts = t.Amounts
		var newHash = sha256.New()
		newHash.Write([]byte(Account.Adress))
		uto := transaction.UnspentTxOut{hex.EncodeToString(newHash.Sum(nil)), len(BlockChain.Chain)-1, Account.Adress, Account.Account}
		_, privateKey := transaction.GetKeys(t.Name)
		tAction := transaction.NewTransaction(privateKey, uto, Account.Adresses, t.Amounts)
		BlockChain.AddBlock(tAction)
		fmt.Println(string(BlockChain.GetJSONBlockChain()))
		ioutil.WriteFile("test.json", BlockChain.GetJSONBlockChain(), 0644)
		fmt.Fprintf(w,"Success")
	}else{
		fmt.Fprintf(w,"error")
	}

}

func saveBlockChain(b block.BlockChain){
	ioutil.WriteFile("test.json", b.GetJSONBlockChain(), 0644)
}

func getBlockChain() []byte{
	byteValue, _ := ioutil.ReadFile("test.json")
	var blockChain block.BlockChain
	json.Unmarshal(byteValue, &blockChain)
	return blockChain.GetJSONBlockChain()
}