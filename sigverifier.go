package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"hash"
	"os"
	"sync"
	"unsafe"
)

var secretKey []byte

func init() {
	if k, e := os.LookupEnv("PLEROMA_SECRET_KEY_BASE"); e {
		secretKey = []byte(k)
	}
}

var hmacPool = sync.Pool{
	New: func() any {
		return hmac.New(sha1.New, secretKey)
	},
}

func verifySig64(base64str string, sigBytes []byte) bool {
	if len(secretKey) < 1 {
		// user didn't put secret key.
		return true
	}

	h := hmacPool.Get().(hash.Hash)
	h.Reset()

	defer hmacPool.Put(h)

	// wanna see a sin?
	// given that string is immutable, HOW ABOUT WE ACCESS THE BYTES, DIRECTLY
	b := unsafe.Slice(unsafe.StringData(base64str), len(base64str))
	h.Write(b)

	var macBuf [sha1.Size]byte
	expectedSig := h.Sum(macBuf[:0])

	return hmac.Equal(sigBytes, expectedSig)
}
