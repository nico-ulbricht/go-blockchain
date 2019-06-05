package chain

import "time"

// Block .
type Block interface {
	GetIndex() uint
	GetPreviousHash() string
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

// New .
func New(previousBlock Block, challengeSolution uint) Block {
	return &block{
		ChallengeSolution: challengeSolution,
		Index:             previousBlock.GetIndex() + 1,
		PreviousHash:      previousBlock.GetPreviousHash(),
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
