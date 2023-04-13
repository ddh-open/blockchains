package blk

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
	"time"
)

const dbName = "blockchain.db"

const blockTableName = "block"

type BlockChain struct {
	Tips []byte
	DB   *bolt.DB
}

// Iterator 创建一个迭代器
func (bc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{bc.Tips, bc.DB}
}

// PrintChain 遍历输出所有区块的信息
func (bc *BlockChain) PrintChain() {
	blockchainIterator := bc.Iterator()
	for {
		block := blockchainIterator.Next()
		fmt.Printf("Height：%d\n", block.Height)
		fmt.Printf("PrevBlockHash：%x\n", block.PrevBlockHash)
		fmt.Printf("Data：%s\n", block.Data)
		fmt.Printf("Timestamp：%s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash：%x\n", block.Hash)
		fmt.Printf("Nonce：%d\n", block.Nonce)
		fmt.Println()
		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)
		// Cmp compares x and y and returns:
		//
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
	}
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

// DBExists 判断数据库是否存在
func DBExists() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}

	return true
}

// CreateBlockChainWithGenesisBlock 创世区块
func CreateBlockChainWithGenesisBlock(data string) *BlockChain {
	// 判断数据库是否存在
	if DBExists() {
		fmt.Println("创世区块已经存在.......")
		os.Exit(1)
	}

	fmt.Println("正在创建创世区块.......")
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
			genesisBlock := CreateGenesisBlock(data)
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

// BlockchainObject 返回Blockchain对象
func BlockchainObject() *BlockChain {

	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var tip []byte
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			// 读取最新区块的Hash
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	return &BlockChain{tip, db}
}
