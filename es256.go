package macsignbench

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

func GenerateSecp256R1Keys() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

func SignSha256(key *ecdsa.PrivateKey, buffer []byte) (r, s *big.Int, err error) {
	h := sha256.New()
	hash := h.Sum(buffer)
	return ecdsa.Sign(rand.Reader, key, hash)
}

func VerifySha256(key *ecdsa.PublicKey, r, s *big.Int, buffer []byte) bool {
	h := sha256.New()
	hash := h.Sum(buffer)
	return ecdsa.Verify(key, hash, r, s)
}
