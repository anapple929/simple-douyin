package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

//Sha256加密
func Sha256(src string) string {
	m := sha256.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
