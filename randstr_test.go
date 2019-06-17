package main

import (
	crand "crypto/rand"
	"encoding/base64"
	"math/rand"
	"testing"

	xrand "golang.org/x/exp/rand"
)

const (
	// Taken from base64 encodeURL alphabet
	characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
	length     = int32(len(characters))
)

func newRandString(n int) string {
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = characters[rand.Int31n(length)]
	}
	return string(buf)
}

func newRandStringBuf(n int) string {
	buf := make([]byte, n)
	rand.Read(buf)

	// Transform in-place
	for i, _ := range buf {
		buf[i] = characters[int32(buf[i])%length]
	}

	return string(buf)
}

func newRandStringBase64(n int) string {
	buf := make([]byte, n)
	rand.Read(buf)

	return base64.RawURLEncoding.EncodeToString(buf)
}

func newCRandStringBuf(n int) string {
	buf := make([]byte, n)
	crand.Read(buf)

	// Transform in-place
	for i, _ := range buf {
		buf[i] = characters[int32(buf[i])%length]
	}

	return string(buf)
}

func newXRandStringBuf(n int) string {
	buf := make([]byte, n)
	xrand.Read(buf)

	// Transform in-place
	for i, _ := range buf {
		buf[i] = characters[int32(buf[i])%length]
	}

	return string(buf)
}

func BenchmarkNaive8(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newRandString(8)
	}
}

func BenchmarkNaive16(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newRandString(16)
	}
}

func BenchmarkNaive32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newRandString(32)
	}
}

func BenchmarkBuf8(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newRandStringBuf(8)
	}
}

func BenchmarkBuf16(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newRandStringBuf(16)
	}
}

func BenchmarkBuf32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newRandStringBuf(32)
	}
}

func BenchmarkBase648(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newRandStringBase64(8)
	}
}

func BenchmarkBase6416(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newRandStringBase64(16)
	}
}

func BenchmarkBase6432(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newRandStringBase64(32)
	}
}

func BenchmarkCrandBuf8(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newCRandStringBuf(8)
	}
}

func BenchmarkCrandBuf16(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newCRandStringBuf(16)
	}
}

func BenchmarkCrandBuf32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newCRandStringBuf(32)
	}
}

func BenchmarkXrandBuf8(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newXRandStringBuf(8)
	}
}

func BenchmarkXrandBuf16(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newXRandStringBuf(16)
	}
}

func BenchmarkXrandBuf32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newXRandStringBuf(32)
	}
}
