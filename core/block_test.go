package core

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/NikhilSharmaWe/go-blockchain/types"
	"github.com/stretchr/testify/assert"
)

func TestHeader_Encode_Decode(t *testing.T) {
	h := Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     2345235,
	}
	fmt.Printf("%+v\n", h)

	buf := &bytes.Buffer{}

	assert.Nil(t, h.EncodeBinary(buf))

	hDecode := Header{}

	assert.Nil(t, hDecode.DecodeBinary(buf))

	assert.Equal(t, h, hDecode)

	fmt.Printf("%+v\n", hDecode)
}

func TestBlock_Encode_Decode(t *testing.T) {
	b := Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height:    10,
			Nonce:     2345235,
		},
		Transactions: nil,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buf))

	bDecode := Block{}

	assert.Nil(t, bDecode.DecodeBinary(buf))
	assert.Equal(t, b, bDecode)

	fmt.Printf("%+v\n", bDecode)
}

func TestBlockHash(t *testing.T) {
	b := Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height:    10,
			Nonce:     2345235,
		},
		Transactions: nil,
	}

	b.Hash()
	fmt.Println(b.hash)
	assert.Equal(t, false, b.hash.IsZero())
}
