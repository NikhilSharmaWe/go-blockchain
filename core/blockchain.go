package core

type Blockchain struct {
	store     Storage
	headers   []*Header
	validator Validator
}

// A genesis block is the first block in a blockchain. It is often hardcoded into the blockchain's software and serves as the foundation upon which the entire blockchain is built. The genesis block typically has special characteristics compared to subsequent blocks in the chain
func NewBlockchain(genesis *Block) (*Blockchain, error) {
	bc := &Blockchain{
		headers: make([]*Header, 0),
		store:   NewMemoryStore(),
	}

	bc.validator = NewBlockValidator(bc)

	if err := bc.addBlockWithoutValidation(genesis); err != nil {
		return nil, err
	}

	return bc, nil
}

func (bc *Blockchain) SetValidator(v Validator) {
	bc.validator = v
}

func (bc *Blockchain) AddBlock(b *Block) error {
	//validate
	if err := bc.validator.Validate(b); err != nil {
		return err
	}

	return bc.addBlockWithoutValidation(b)
}

func (bc *Blockchain) HasBlock(height uint32) bool {
	return height <= bc.Height()
}

// [0, 1, 2 ,3] => 4 len
// [0, 1, 2 ,3] => 3 height
func (bc *Blockchain) Height() uint32 {
	return uint32(len(bc.headers) - 1)
}

func (bc *Blockchain) addBlockWithoutValidation(b *Block) error {
	bc.headers = append(bc.headers, b.Header)
	return bc.store.Put(b)
}
