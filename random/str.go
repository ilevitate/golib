package random

import (
	"bytes"
	"math/rand"
	"time"
	"unsafe"
)

//指定长度的随机大写字母
func RandUppercase(l int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < l; {
		if string(RandInt(65, 90)) != temp {
			temp = string(RandInt(65, 90))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}

//指定长度的随机小写字母
func RandLowercase(l int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < l; {
		if string(RandInt(97, 122)) != temp {
			temp = string(RandInt(97, 122))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}

//指定长度的随机字符串，第二个参数限制只能出现指定的字符
//指定的字符里不兼容汉字
func RandStr(l int, specifiedStr ...string) string {
	var tpl string
	if len(specifiedStr) > 0 {
		tpl = specifiedStr[0]
	} else {
		tpl = "abcdefghijklmnopqrstuwxyzABCDEFGHIJKLMNOPQRSTUWXYZ0123456789"
	}
	tplLen := len(tpl)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, l)
	for i := 0; i < l; i++ {
		b[i] = tpl[r.Intn(tplLen)]
	}
	return bytesToStr(b)
}

func bytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
