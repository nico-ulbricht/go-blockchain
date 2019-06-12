package chain

// Chain .
type Chain interface {
	AddBlock(b Block) Block
	IsChainValid() bool
	GetLength() int
}

type chain []Block

func (c *chain) AddBlock(b Block) Block {
	addedChain := append(*c, b)
	*c = addedChain
	return b
}

func (c *chain) GetLength() int {
	return len(*c)
}

func (c *chain) IsChainValid() bool {
	for _, aBlock := range *c {
		if aBlock.IsSolutionValid() == false {
			return false
		}
	}

	return true
}

func NewChain() Chain {
	return &chain{}
}
