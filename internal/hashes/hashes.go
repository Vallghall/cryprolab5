package hashes

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"hash"

	"github.com/ddulesov/pkcs7"
	"github.com/maoxs2/go-ripemd"
)

var hashes map[string]hash.Hash

func init() {
	hashes = map[string]hash.Hash{
		"SHA-224":    sha256.New224(),
		"SHA-256":    sha256.New(),
		"SHA-384":    sha512.New384(),
		"SHA-512":    sha512.New(),
		"RIPEMD-128": ripemd.New128(),
		"RIPEMD-160": ripemd.New160(),
		"RIPEMD-256": ripemd.New256(),
		"RIPEMD-320": ripemd.New320(),
		"Md5":        md5.New(),
		"GOST":       pkcs7.GOSTR3411_94.New(),
	}
}
