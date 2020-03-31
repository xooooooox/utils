package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 md5加密
func Md5(bs []byte) string {
	hash := md5.New()
	hash.Write(bs)
	return hex.EncodeToString(hash.Sum(nil))
}
