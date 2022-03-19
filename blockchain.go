package main

import (
	"fmt"
	"time"
	"crypto/sha256"
	"strings"
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
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v",b)))
	return fmt.Sprintf("%x",h.Sum(nil))
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
}

func (b Blockchain) isValid() bool {
	for i := range b.chain[1:] {
		pblock := &b.chain[i]
		block := &b.chain[i+1]
		if block.hash != block.hashblock() || block.hash_p != pblock.hash {
			fmt.Println(block.hash)
			fmt.Println(block.hashblock())
			fmt.Println(block.hash_p)
			fmt.Println(pblock.hash)
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



