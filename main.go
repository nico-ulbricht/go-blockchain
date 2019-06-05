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

	challenge := uint(0)
	for {
		block := chain.New(previousBlock, challenge)
		isValid := block.IsSolutionValid()
		fmt.Println(challenge, block.GetHash(), isValid)
		if isValid == true {
			previousBlock = block
			challenge = 0
			time.Sleep(time.Second * 2)
		}

		challenge++
	}
}
