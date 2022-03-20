package main

import (
	"fmt"
	"time"
	"crypto/sha256"
	"strings"
	"encoding/json"
	"strconv"
)

type Block struct {
	data string
	hash string
	hash_p string
	timestamp time.Time
	pow int
}

type Blockchain struct {
	headblock Block
	chain []Block
	diff int
}

func (b *Block)hashblock() string {
    data, _ := json.Marshal(b.data)
    blockData := b.hash_p + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
    blockHash := sha256.Sum256([]byte(blockData))
    return fmt.Sprintf("%x", blockHash)
}

func (b *Block) mine(diff int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0",diff)) {
		b.pow ++
		b.hash = b.hashblock() 
	}
}

func initBlockChain(diff int) Blockchain {
	headblock := Block{
		hash : "0",
		timestamp : time.Now(),
	}

	return Blockchain {
		headblock,
		[]Block{headblock},
		diff,
	}
}

func (b *Blockchain) addblock(data string) {
	lastblock := b.chain[len(b.chain)-1]
	newblock := Block{
		data : data,
		hash_p : lastblock.hash,
		timestamp : time.Now(),
	}
	newblock.mine(b.diff)
	b.chain = append(b.chain, newblock)
	fmt.Printf("-----------------------------\nAdded new block to blockchain.\nData :%s\nHash of previous : %s\nHash : %s\n-----------------------------\n", newblock.data, newblock.hash_p, newblock.hash)
}

func (b Blockchain) isValid() bool {
	for i := range b.chain[1:] {
		pblock := &b.chain[i]
		block := &b.chain[i+1]
		if block.hash != block.hashblock() || block.hash_p != pblock.hash {
			fmt.Printf("Checking block %d\n",i)
			fmt.Printf("Block %d hash (from block.hash) : %s\n",i,block.hash)
			fmt.Printf("Block %d hash (from block.hashblock()) : %s\n",i,block.hashblock())
			fmt.Printf("Block %d : previous block hash (from block.hash_p) : %s\n",i, block.hash_p)
			fmt.Printf("Block %d : previous block hash (from pblock.hash) : %s\n",i, pblock.hash)
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("Blockchain started at %s\n",time.Now().String())
	blockchain := initBlockChain(2)
	blockchain.addblock("Pain")
	blockchain.addblock("Oeufs")
	fmt.Println(blockchain.isValid())	
}
