package chain

import (
	"time"
	"crypto/sha256"
)

// Block ...
type Block struct {
	hash []byte
	Content string
	timestamp time.Time 
	previousHash []byte
}

// NewBlock ...
func NewBlock(content string, previousHash []byte) Block {
	hasher := sha256.New();
	timestamp := time.Now();
	
	splitedContent := timestamp.String() + string(previousHash) + content
	hasher.Write([]byte(splitedContent))

	return Block {
		timestamp: time.Now(),
		hash: hasher.Sum(nil),
		Content: content,
		previousHash: previousHash,
	}
}