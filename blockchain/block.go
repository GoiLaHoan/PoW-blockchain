package blockchain

import (
    "time"
)

type BlockChain struct {
    Blocks []*Block
}

type Block struct {
    TimeStamp int64
    Hash     []byte
    Data     []byte
    PrevHash []byte
    Nonce    int
    
}

func CreateBlock(data string, prevHash []byte) *Block {
    block := &Block{time.Now().Unix(), []byte{}, []byte(data), prevHash, 0} 
        // Don't forget to add the 0 at the end for the nonce!
    pow := NewProofOfWork(block)
    nonce, hash := pow.Run()

    block.Hash = hash[:]
    block.Nonce = nonce

    return block
}

func (chain *BlockChain) AddBlock(data string) {
    prevBlock := chain.Blocks[len(chain.Blocks)-1]
    new := CreateBlock(data, prevBlock.Hash)
    chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis Block", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}