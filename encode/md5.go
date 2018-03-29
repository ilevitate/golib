package encode

import (
	"crypto/md5"
	"encoding/hex"
	"unsafe"
)

//从[]byte生成md5密文
func MD5ByBytes(value []byte, salt ...[]byte) (result string, err error) {
	var s []byte
	if len(salt) > 0 {
		s = salt[0]
	}
	h := md5.New()
	_, err = h.Write(value)
	if err != nil {
		return
	}
	if s != nil {
		_, err = h.Write(s)
		if err != nil {
			return
		}
	}
	result = hex.EncodeToString(h.Sum(nil))
	return
}

//从str生成md5密文
func MD5ByStr(value string, salt ...string) (result string, err error) {
	var s []byte
	if len(salt) > 0 {
		s = strToByte(salt[0])
	}
	h := md5.New()
	_, err = h.Write(strToByte(value))
	if err != nil {
		return
	}
	if s != nil {
		_, err = h.Write(s)
		if err != nil {
			return
		}
	}
	result = hex.EncodeToString(h.Sum(nil))
	return
}

func strToByte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
