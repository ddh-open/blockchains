package main

import (
	"blockchains/blk"
	"fmt"
)

func main() {
	// 创世区块
	blockchain := blk.CreateBlockChainWithGenesisBlock()
	// 新区块
	blockchain.AddBlockToBlockChain("Send 100RMB To zhangqiang", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	blockchain.AddBlockToBlockChain("Send 200RMB To changjingkong", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	blockchain.AddBlockToBlockChain("Send 300RMB To juncheng", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	blockchain.AddBlockToBlockChain("Send 50RMB To haolin", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	fmt.Println(blockchain)
	fmt.Println(blockchain.Blocks)
}
