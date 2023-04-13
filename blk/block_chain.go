package blk

import (
	"github.com/boltdb/bolt"
	"log"
)

const dbName = "blockchain.db"

const blockTableName = "block"

type BlockChain struct {
	Tips []byte
	DB   *bolt.DB
}

// AddBlockToBlockChain 添加区块
func (bc *BlockChain) AddBlockToBlockChain(data string) {
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		// 获取表
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			blockBytes := b.Get(bc.Tips)
			block := Deserialize(blockBytes)

			newBlock := NewBlock(data, block.Height+1, block.Hash)
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			bc.Tips = newBlock.Hash
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

// CreateBlockChainWithGenesisBlock 创世区块
func CreateBlockChainWithGenesisBlock() *BlockChain {
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var blockHash []byte

	err = db.Update(func(tx *bolt.Tx) error {
		// 获取表
		b := tx.Bucket([]byte(blockTableName))

		if b == nil {
			b, err = tx.CreateBucket([]byte(blockTableName))
			if err != nil {
				log.Fatal(err)
			}
		}

		if b != nil {
			genesisBlock := CreateGenesisBlock("Genesis Block......")
			err = b.Put(genesisBlock.Hash, genesisBlock.Serialize())

			if err != nil {
				log.Panic(err)
			}
			blockHash = genesisBlock.Hash
		}

		return nil
	})

	return &BlockChain{blockHash, db}
}
