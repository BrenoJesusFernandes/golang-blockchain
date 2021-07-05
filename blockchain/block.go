package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	// Criando hash baseado nos dados e no hash anterior
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)

	//Add hash no bloco
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()

	return block
}

func (chain *BlockChain) AddBlock(data string) {
	//Localiza bloco anterior
	prevBlock := chain.Blocks[len(chain.Blocks)-1]

	//Cria o novo bloco
	newBlock := CreateBlock(data, prevBlock.Hash)

	// Add o novo bloco na corrent
	chain.Blocks = append(chain.Blocks, newBlock)
}

// Genesis Cria o primeiro bloco
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain Inicia a blockchain
func InitBlockChain() *BlockChain {
	GenesisBlock := Genesis()
	return &BlockChain{[]*Block{GenesisBlock}}
}
