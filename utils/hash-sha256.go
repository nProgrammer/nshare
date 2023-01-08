package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(text string) string {
	bv := []byte(text)
	hash := sha256.Sum256(bv)
	sha := hex.EncodeToString(hash[:])
	return sha
}
