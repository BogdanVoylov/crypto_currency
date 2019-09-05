package block

import (
	"crypto/sha256"
	"cryptocurrency/transaction"
	"encoding/hex"
	"io"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Index   int
	Hash    string
	PreviousHash string
	Timestamp int64
	Data transaction.Transaction

}

func NewBlock(index int, hash string, previousHash string, timestamp time.Time, data transaction.Transaction) Block {
	return Block{Index: index, Hash: strings.ReplaceAll(hash, "\n","h"), PreviousHash: previousHash, Timestamp: timestamp.Unix(), Data: data}
}
 /*index + previousHash + timestamp + data*/

func calculateHash(index int, previousHash string, timestamp int64, data transaction.Transaction) string {
	var foo = strconv.Itoa(index) + string(previousHash) + strconv.Itoa(int(timestamp)) + hex.EncodeToString(data.Id)
	input :=strings.NewReader(foo)
	var newHash = sha256.New()
	if _, err := io.Copy(newHash, input); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(newHash.Sum(nil))
}

func (previousBlock Block) GenerateNextBlock(data transaction.Transaction) Block{
	index := previousBlock.Index+1
	previousHash := previousBlock.Hash
	t := time.Now()
	timestamp := t.Unix()
	newHash := calculateHash(index, previousHash, timestamp, data)
	return NewBlock(index, newHash, previousHash, time.Now(), data)

}

func RandStringBytes(n int) []byte{
	const letterBytes = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

func generateGenesisBlock() Block{
	input := RandStringBytes(10)
	/*fmt.Println(input)*/
	newHash := sha256.Sum256(input)
	var buff []byte = newHash[0:31]
	return NewBlock(0, hex.EncodeToString(buff),  "",  time.Now(),  transaction.Transaction{nil,nil,nil})
}


/*func isValidNewBlock(previousBlock, newBlock Block) bool{
	if(previousBlock.Index == newBlock.Index-1 && previousBlock.Hash == newBlock.PreviousHash){
		return true
	}else{
		return false
	}
}

func IsValidBlockStructure(block Block) bool{
	if(fmt.Sprintf("%T", block.Index) == "int64" && fmt.Sprintf("%T", block.Hash) == "string" && fmt.Sprintf("%T", block.PreviousHash) == "string" && fmt.Sprintf("%T", block.Timestamp) == "int64" && fmt.Sprintf("%T", block.Data) == "string"){
		return true
	}else{
		return false
	}
}*/