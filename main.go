package main

import (
    "fmt"
    "strconv"
    "github.com/GoiLaHoan/blockchain/blockchain"
)

func main() {

    chain := blockchain.InitBlockChain()

    chain.AddBlock("Send 9 BTC to Hoan")
    chain.AddBlock("Send 11 BTC to Long")

    for _, block := range chain.Blocks {

        fmt.Printf("TimeStamp: %x\n", block.TimeStamp)
        fmt.Printf("Previous hash: %x\n", block.PrevHash)
        fmt.Printf("data: %s\n", block.Data)
        fmt.Printf("hash: %x\n", block.Hash)

        pow := blockchain.NewProofOfWork(block)
        fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
        fmt.Println()
    }

}
