package blk

type BlockChain struct {
	Blocks []*Block
}

// AddBlockToBlockChain 添加区块
func (bc *BlockChain) AddBlockToBlockChain(data string, height int64, prevBlockHash []byte) {
	block := NewBlock(data, height, prevBlockHash)
	bc.Blocks = append(bc.Blocks, block)
}

// CreateBlockChainWithGenesisBlock 创世区块
func CreateBlockChainWithGenesisBlock() *BlockChain {
	genesisBlock := CreateGenesisBlock("Genesis Block......")
	return &BlockChain{
		Blocks: append([]*Block{}, genesisBlock),
	}
}
