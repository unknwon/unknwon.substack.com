package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"testing"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
)

var (
	secret = []byte("Hello, World!")
	salt   = []byte("randomsalt123456") // 16-byte salt
)

func BenchmarkSHA256(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SHA256(secret)
	}
}

func BenchmarkSHA512(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SHA512(secret)
	}
}

func BenchmarkArgon2(b *testing.B) {
	time := uint32(1)
	memory := uint32(64 * 1024)
	threads := uint8(4)
	keyLen := uint32(32)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		argon2.IDKey(secret, salt, time, memory, threads, keyLen)
	}
}

func BenchmarkBcrypt(b *testing.B) {
	cost := 10 // Standard bcrypt cost

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bcrypt.GenerateFromPassword(secret, cost)
	}
}

func BenchmarkScrypt(b *testing.B) {
	N := 16384 // CPU/memory cost parameter (power of 2)
	r := 8     // block size parameter
	p := 1     // parallelization parameter
	keyLen := 32

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		scrypt.Key(secret, salt, N, r, p, keyLen)
	}
}

func BenchmarkPBKDF2_SHA256(b *testing.B) {
	iterations := 210000 // OWASP recommended minimum
	keyLen := 32

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pbkdf2.Key(secret, salt, iterations, keyLen, sha256.New)
	}
}

func BenchmarkPBKDF2_SHA512(b *testing.B) {
	iterations := 210000 // OWASP recommended minimum
	keyLen := 32

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pbkdf2.Key(secret, salt, iterations, keyLen, sha512.New)
	}
}
