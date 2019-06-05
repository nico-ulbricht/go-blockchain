package chain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// Block .
type Block interface {
	GetHash() string
	GetIndex() uint
	GetPreviousHash() string
	IsSolutionValid() bool
}

type block struct {
	ChallengeSolution uint
	Index             uint
	PreviousHash      string
	Timestamp         time.Time
}

func (b *block) GetIndex() uint {
	return b.Index
}

func (b *block) GetPreviousHash() string {
	return b.PreviousHash
}

func (b *block) GetHash() string {
	data := fmt.Sprintf("%s%d", b.PreviousHash, b.ChallengeSolution)
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

func (b *block) IsSolutionValid() bool {
	hash := b.GetHash()
	return strings.HasSuffix(hash, "0")
}

// New .
func New(previousBlock Block, challengeSolution uint) Block {
	return &block{
		ChallengeSolution: challengeSolution,
		Index:             previousBlock.GetIndex() + 1,
		PreviousHash:      previousBlock.GetHash(),
		Timestamp:         time.Now(),
	}
}

// NewGenesis .
func NewGenesis() Block {
	return &block{
		ChallengeSolution: 100,
		Index:             0,
		PreviousHash:      "0",
		Timestamp:         time.Now(),
	}
}
