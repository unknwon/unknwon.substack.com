package main

import (
	"crypto/sha256"
	"crypto/sha512"
)

func SHA256(data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func SHA512(data []byte) []byte {
	hash := sha512.New()
	hash.Write(data)
	return hash.Sum(nil)
}
