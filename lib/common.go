package lib

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {
	var m = md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}

func Uint8sToBytes(src []uint8) []byte {
	var dst []byte
	for _, b := range src {
		dst = append(dst, byte(b))
	}
	return dst
}
