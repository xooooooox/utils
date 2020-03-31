package utils

import (
	"math/rand"
)

// RandString([]byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}, 8)
// RandString([]byte("1234567890"), 8)
// RandString 随机字符串
func RandString(bs []byte, length int) string {
	bytesLength := len(bs)
	if bytesLength == 0 {
		bs = []byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
		bytesLength = len(bs)
	}
	if length < 1 {
		length = 6
	}
	result := make([]byte, length, length)
	var tmp byte
	for i := 0; i < length; i++ {
		tmp = bs[rand.Intn(bytesLength)]
		result = append(result, tmp)
	}
	return string(result[:])
}
