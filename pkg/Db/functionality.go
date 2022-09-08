package Db

import (
	"crypto/md5"
	"encoding/hex"
)

func GetHash(input string) string {
	h := md5.New()
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}
