package macsignbench

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"hash"
)

var ErrFailedToGetErrorBuffer = errors.New("")

func NewHmacSha256(key [32]byte) hash.Hash {
	return hmac.New(sha256.New, key[:])
}

func GenerateHmacKey() ([32]byte, error) {
	randBuffer := make([]byte, 32)
	count, err := rand.Read(randBuffer)
	if err != nil {
		return [32]byte{}, nil
	}
	if count != 32 {
		return [32]byte{}, fmt.Errorf("expected to have read 32 random bytes, but instead read %d", count)
	}
	var key [32]byte
	count = copy(key[:], randBuffer)
	if count != 32 {
		return key, fmt.Errorf("expected to have copied 32 bytes, but instead read %d", count)
	}
	return key, nil
}
