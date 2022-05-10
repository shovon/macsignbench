package macsignbench

import (
	"crypto/hmac"
	"crypto/rand"
	"errors"
	"fmt"
	"testing"
)

const payloadLength = 128

func BenchmarkHmacSha256(b *testing.B) {
	key, err := GenerateHmacKey()
	if err != nil {
		panic(err)
	}
	hasher := NewHmacSha256(key)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		payload := make([]byte, payloadLength)
		n, err := rand.Read(payload)
		if err != nil {
			panic(err)
		}
		if n != payloadLength {
			panic(fmt.Errorf("expected to have read %d bytes, but instead got %d", payloadLength, n))
		}
		b.StartTimer()
		result1 := hasher.Sum(payload)
		result2 := hasher.Sum(payload)
		if !hmac.Equal(result1, result2) {
			b.Error(errors.New("expected the hashes to match, but they don't"))
			b.FailNow()
		}
	}
}
