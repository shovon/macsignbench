package macsignbench

import (
	"crypto/ed25519"
	"crypto/rand"
	"errors"
	"fmt"
	"testing"
)

func BenchmarkEd25519(b *testing.B) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

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
		signature := ed25519.Sign(privateKey, payload)
		if !ed25519.Verify(publicKey, payload, signature) {
			b.Error(errors.New("expected the signature to validate, but it didn't"))
			b.FailNow()
		}
	}
}
