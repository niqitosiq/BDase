package chain

import (
	"log"
)

// Chain ...
type Chain struct {
	Name string
	Blocks []Block
}

var chains []*Chain

// NewChain ...
func NewChain(name string) *Chain {
	chain := &Chain {
		Name: name,
		Blocks: []Block {},
	}

	chains = append(chains, chain)

	return chain;
}

// AppendBlock ...
func AppendBlock(name string, content string) *Chain {
	var currentChain *Chain

	for _, chain := range chains {
		if (chain.Name == name) {
			currentChain = chain
		}
	}
	if currentChain == nil {
		log.Print("Цепь не найдена")
		return currentChain
	}

	blocks := currentChain.Blocks

	lastIndex := len(blocks) - 1

	var lastHash []byte
	if lastIndex == -1 {
		lastHash = []byte("")
	} else {
		lastHash = blocks[lastIndex].hash
	}
	
	block := NewBlock(content, lastHash);

	currentChain.Blocks = append(blocks, block)
	return currentChain
}