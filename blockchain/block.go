package blockchain

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)

	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

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
