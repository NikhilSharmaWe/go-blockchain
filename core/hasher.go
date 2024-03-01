package core

import (
	"crypto/sha256"

	"github.com/NikhilSharmaWe/go-blockchain/types"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}

type BlockHasher struct{}

func (BlockHasher) Hash(b *Block) types.Hash {
	return sha256.Sum256(b.HeaderData())
}
