package blk

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockChainIterator struct {
	CurrentHash []byte
	DB          *bolt.DB
}

func (bci *BlockChainIterator) Next() *Block {
	var block *Block

	err := bci.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			blockBytes := b.Get(bci.CurrentHash)
			block = Deserialize(blockBytes)
			bci.CurrentHash = block.PrevBlockHash
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	return block
}
