package block

import (
	"cryptocurrency/transaction"
	"encoding/json"
	"fmt"
)

type BlockChain struct {
	Chain []Block
}

func NewBlockChain() *BlockChain {
	return &BlockChain{Chain: []Block{generateGenesisBlock()}}
}

func (b *BlockChain) AddBlock(data transaction.Transaction){
	fmt.Println(len(b.Chain))
	block := b.Chain[len(b.Chain)-1]
	b.Chain = append(b.Chain, block.GenerateNextBlock(data))
}

func (b BlockChain) GetBlockChain() *BlockChain {
	return &b
}

func (b BlockChain) GetJSONBlockChain() []byte{
	result, _ := json.Marshal(b)
	fmt.Println(string(result))
	return result
}