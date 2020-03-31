package utils

import (
	"encoding/base64"
)

// Base64Encrypt base64加密
func Base64Encrypt(bs []byte) string {
	return base64.StdEncoding.EncodeToString(bs)
}

// Base64Decrypt base64解密
func Base64Decrypt(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
