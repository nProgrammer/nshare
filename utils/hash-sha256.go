package utils

import "crypto/sha256"

func Sha256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return string(h.Sum(nil))
}
