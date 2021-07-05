package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	blocks []*Block
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
	prevBlock := chain.blocks[len(chain.blocks)-1]

	//Cria o novo bloco
	newBlock := CreateBlock(data, prevBlock.Hash)

	// Add o novo bloco na corrent
	chain.blocks = append(chain.blocks, newBlock)
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

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x \n", block.PrevHash)
		fmt.Printf("Data in Block: %s \n", block.Data)
		fmt.Printf("Hash: %x \n", block.Hash)

		fmt.Printf("---------------\n")
	}
}
