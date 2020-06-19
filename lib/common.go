package lib

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"runtime"
	"strconv"
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

func AssetsError(err error) {
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		buffer := new(bytes.Buffer)
		if ok {
			buffer.WriteString("file: ")
			buffer.WriteString(file)
			buffer.WriteString(" line: ")
			buffer.WriteString(strconv.Itoa(line))
			buffer.WriteString(" err: ")
		}
		buffer.WriteString(err.Error())
		panic(errors.New(buffer.String()))
	}
}
