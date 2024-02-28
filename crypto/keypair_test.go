package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	fmt.Printf("%+v\n", privKey.key)

	pubKey := privKey.PublicKey()
	fmt.Printf("%+v\n", pubKey.key)

	address := pubKey.Address()
	fmt.Println(address)

	msg := []byte("hello")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	fmt.Printf("%+v\n", sig)

	assert.True(t, sig.Verify(pubKey, msg))
}

func TestKeyPairSignVerifySuccess(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	msg := []byte("hello")

	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	assert.True(t, sig.Verify(pubKey, msg))
}

func TestKeyPairSignVerifyFail(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	msg := []byte("hello")

	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	otherPrivKey := GeneratePrivateKey()
	otherPubKey := otherPrivKey.PublicKey()
	otherMsg := []byte("other message")

	assert.False(t, sig.Verify(otherPubKey, msg))
	assert.False(t, sig.Verify(pubKey, otherMsg))
}
