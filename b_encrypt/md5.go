package b_encrypt

import (
	"crypto/md5"
	"fmt"
)

func Md5Hex(dataStr string) string {
	data := []byte(dataStr)
	md5Sum := md5.Sum(data)
	return fmt.Sprintf("%x", md5Sum) //将[]byte转成16进制
}
