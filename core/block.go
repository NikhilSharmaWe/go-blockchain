package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/NikhilSharmaWe/go-blockchain/crypto"
	"github.com/NikhilSharmaWe/go-blockchain/types"
)

// why we need header in a block?
// if we want to hash the block we cannot hash everything
// for saving space: if we want to keep track of blocks but we do not want the transactions for that
type Header struct {
	Version       uint32
	DataHash      types.Hash // Data hash should be present in the Header so that the Header is changed when data is changed of the block
	PrevBlockHash types.Hash
	Timestamp     int64
	Height        uint32
	Nonce         uint64
}

type Block struct {
	*Header      // using *Header is more memory efficient than using Header
	Transactions []Transaction

	// PublicKey and Signature are added in the Block also (Transaction obviously need them) because in case when a new leader is elected and want to change the Block then we can verify that whether should we allow the change or not
	Validator crypto.PublicKey
	Signature *crypto.Signature

	// cached version of the header hash
	hash types.Hash
}

func NewBlock(h *Header, txx []Transaction) *Block {
	return &Block{
		Header:       h,
		Transactions: txx,
	}
}

func (b *Block) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(b.HeaderData())
	if err != nil {
		return err
	}

	b.Validator = privKey.PublicKey()
	b.Signature = sig

	return nil
}

func (b *Block) Verify() error {
	if b.Signature == nil {
		return fmt.Errorf("block has no signature")
	}
	if !b.Signature.Verify(b.Validator, b.HeaderData()) {
		return fmt.Errorf("block has invalid signature")
	}

	return nil
}

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(r, b)
}

func (b *Block) Encoder(w io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(w, b)
}

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b)
	}

	return b.hash
}

func (b *Block) HeaderData() []byte {
	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(b.Header)
	return buf.Bytes()
}
