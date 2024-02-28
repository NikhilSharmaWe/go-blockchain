package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"

	"github.com/NikhilSharmaWe/go-blockchain/types"
)

// Private Key:
// Think of your private key like a super-secret password that only you know. It's a long string of numbers and letters that gives you control over your digital belongings, like cryptocurrencies (think of them as digital money) in the blockchain world.
// Public Key:
// Now, the public key is like your public address or username. It's derived from your private key through some fancy math. You share your public key openly with others so they can send you things, like cryptocurrencies or messages, on the blockchain.
// Signature:
// Imagine you want to sign a letter to prove it's really from you. In the digital world, a signature is like a digital version of your handwritten signature. It's created using your private key and it proves that you're the one who authorized a transaction or message.

// So, here's how it works:

// You have your private key (your secret password) and your public key (your public address).
// When you want to send something, like a cryptocurrency transaction, you use your private key to create a unique signature.
// You send your transaction along with this signature.
// Others can use your public key to check the signature and make sure it's really from you.
// If everything checks out, the transaction is considered valid and gets added to the blockchain.
// In essence, private keys let you securely control your digital assets, public keys identify you on the network, and signatures prove that you authorized transactions. Together, they make blockchain transactions secure and trustworthy!

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

func (k PrivateKey) Sign(data []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, k.key, data)
	if err != nil {
		return nil, err
	}

	return &Signature{
		r: r,
		s: s,
	}, nil
}

func GeneratePrivateKey() PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	return PrivateKey{
		key: key,
	}
}

type PublicKey struct {
	key *ecdsa.PublicKey
}

func (k PrivateKey) PublicKey() PublicKey {
	return PublicKey{
		key: &k.key.PublicKey,
	}
}

func (k PublicKey) ToSlice() []byte {
	return elliptic.MarshalCompressed(k.key, k.key.X, k.key.Y)
}

func (k *PublicKey) Address() types.Address {
	h := sha256.Sum256(k.ToSlice())
	return types.AddressFromBytes(h[len(h)-20:])
}

type Signature struct {
	r, s *big.Int
}

func (sig *Signature) Verify(pubKey PublicKey, data []byte) bool {
	return ecdsa.Verify(pubKey.key, data, sig.r, sig.s)
}
