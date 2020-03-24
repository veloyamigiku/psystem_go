package auth

import (
	"crypto/md5"
	"encoding/hex"
)

func FromStringToMD5(str string) (md5Str string) {
	passwordHash := md5.Sum([]byte(str))
	md5Str = hex.EncodeToString(passwordHash[:])
	return
}
