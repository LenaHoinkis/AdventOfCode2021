package utils

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

func HexToBinary(s string) string {
	decoded, _ := hex.DecodeString(s)
	var sb strings.Builder
	for _, bit := range decoded {
		sb.WriteString(fmt.Sprintf("%08b", int(bit)))
	}
	return sb.String()
}

func BinaryToInt(s string) (int64, error) {
	return strconv.ParseInt(clone(s), 2, 64)
}

func clone(s string) string {
	b := make([]byte, len(s))
	copy(b, s)
	return *(*string)(unsafe.Pointer(&b))
}
