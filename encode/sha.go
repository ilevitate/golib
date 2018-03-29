package encode

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"

	"convert"
)

//根据[]byte生成sha1密文
func SHA1ByBytes(value []byte, salt ...[]byte) (result string, err error) {
	var s []byte
	if len(salt) > 0 {
		s = salt[0]
	}
	h := sha1.New()
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

//根据str生成sha1密文
func SHA1ByStr(value string, salt ...string) (result string, err error) {
	var s []byte
	if len(salt) > 0 {
		s = convert.StrToByte(salt[0])
	}
	h := sha1.New()
	_, err = h.Write(convert.StrToByte(value))
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

//根据[]byte生成sha256密文
func SHA256ByBytes(value []byte, salt ...[]byte) (result string, err error) {
	var s []byte
	if len(salt) > 0 {
		s = salt[0]
	}
	h := sha256.New()
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
	//计算出字符串格式的签名
	result = hex.EncodeToString(h.Sum(nil))
	return
}

//根据string生成sha256密文
func SHA256ByStr(value string, salt ...string) (result string, err error) {
	var s []byte
	if len(salt) > 0 {
		s = convert.StrToByte(salt[0])
	}
	h := sha256.New()
	_, err = h.Write(convert.StrToByte(value))
	if err != nil {
		return
	}
	if s != nil {
		_, err = h.Write(s)
		if err != nil {
			return
		}
	}
	//计算出字符串格式的签名
	result = hex.EncodeToString(h.Sum(nil))
	return
}
