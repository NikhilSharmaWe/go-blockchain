package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBlochainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0))
	assert.Nil(t, err)

	return bc
}

func TestNewBlockchain(t *testing.T) {
	bc := newBlochainWithGenesis(t)
	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(0))

	fmt.Println(bc.Height())
}

func TestHasBlock(t *testing.T) {
	bc := newBlochainWithGenesis(t)
	assert.True(t, bc.HasBlock(0))
}

func TestAddBlock(t *testing.T) {
	bc := newBlochainWithGenesis(t)

	lenBlocks := 1000
	for i := 0; i < lenBlocks; i++ {
		block := randomBlockWithSignature(t, uint32(i+1))
		assert.Nil(t, bc.AddBlock(block))
	}

	// len of the blockchain is 10001
	// height of the blockchain is 1000
	assert.Equal(t, uint32(1000), bc.Height())
	assert.Equal(t, 1001, len(bc.headers))
	assert.NotNil(t, bc.AddBlock(randomBlock(89)))
}
