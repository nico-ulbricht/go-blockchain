package chain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Block .
type Block interface {
	GetHash() string
	GetIndex() uint
	IsSolutionValid() bool
}

type block struct {
	ChallengeSolution uint      `json:"challengeSolution"`
	Index             uint      `json:"index"`
	PreviousHash      string    `json:"previousHash"`
	Timestamp         time.Time `json:"timestamp"`
}

func (b *block) GetIndex() uint {
	return b.Index
}

func (b *block) GetHash() string {
	payload, err := json.Marshal(b)
	if err != nil {
		panic(fmt.Errorf("unable to hash block %d, %v", b.Index, err))
	}

	data := fmt.Sprintf("%s-%d", payload, b.ChallengeSolution)
	hash := sha256.New()
	hash.Write([]byte(data))
	hexHash := hex.EncodeToString(hash.Sum(nil))
	return hexHash
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
