package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)
//小写的
func Md5Encode(data string) string{
	h := md5.New()
	h.Write([]byte(data)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return  hex.EncodeToString(cipherStr)
}
//大写
func MD5Encode(data string) string{
	return strings.ToUpper(Md5Encode(data))
}

func ValidatePasswd(plainpwd,salt,passwd string) bool{
	return Md5Encode(plainpwd+salt)==passwd
}
func MakePasswd(plainpwd,salt string) string{
	return Md5Encode(plainpwd+salt)
}
