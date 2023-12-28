package main

import (
	"github.com/xor111xor/blockchain/internal/models"
)

func main() {
	BlockChain := models.NewBlockchain()
	BlockChain.Print()

	previousHash := BlockChain.LastBlock().Hash()
	BlockChain.CreateBlock(2, previousHash)
	BlockChain.Print()
}
