package main

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"github.com/nico-ulbricht/go-blockchain/pkg/chain"
)

func init() {
	godotenv.Load()
}

func main() {
	previousBlock := chain.NewGenesis()
	blockChain := chain.NewChain()

	challenge := uint(0)
	for {
		block := chain.New(previousBlock, challenge)
		isValid := block.IsSolutionValid()
		fmt.Println(challenge, block.GetHash(), isValid)
		if isValid == false {
			challenge++
			continue
		}

		previousBlock = block
		blockChain.AddBlock(block)
		challenge = 0
		if blockChain.GetLength() == 3 {
			break
		}

		time.Sleep(time.Second * 2)
	}

	fmt.Printf("chain validity: %v\n", blockChain.IsChainValid())
}
