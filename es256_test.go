package macsignbench

import (
	"crypto/rand"
	"errors"
	"fmt"
	"testing"
)

func BenchmarkEs256(b *testing.B) {
	key, err := GenerateSecp256R1Keys()
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
		r, s, err := SignSha256(key, payload)
		if err != nil {
			panic(err)
		}
		if !VerifySha256(&key.PublicKey, r, s, payload) {
			b.Error(errors.New("expected the signature to validate, but it didn't"))
			b.FailNow()
		}
	}
}
