//数据校验

package validation

import "regexp"

//判断手机号
func Mobile(str string) bool {
	rule := "^1[34578][0-9]{9}$"
	return regexp.MustCompile(rule).MatchString(str)
}

//判断QQ
func QQ(str string) bool {
	rule := "^[0-9]{5,15}$"
	return regexp.MustCompile(rule).MatchString(str)
}

//判断用户名
//可以包含a-z|0-9
func UserName(str string) bool {
	rule := "^[a-z0-9]*$"
	return regexp.MustCompile(rule).MatchString(str)
}

//判断昵称
//可以包含a-z|A-Z|0-9|_|汉字
//不能以_开头或结尾
func NickName(str string) bool {
	rule := `^[a-z0-9A-Z\p{Han}]+(_[a-z0-9A-Z\p{Han}]+)*$`
	return regexp.MustCompile(rule).MatchString(str)
}

//判断邮箱
//可以包含a-z|0-9|_|-|.
//不能以_-.开头或结尾
func Mail(str string) bool {
	rule := `^[a-z0-9]+([\-_\.][a-z0-9]+)*@([a-z0-9]+(-[a-z0-9]+)*\.)+[a-z]{2,4}$`
	return regexp.MustCompile(rule).MatchString(str)
}

//判断IP
func IP(str string) bool {
	rule := "^((2[0-4]\\d|25[0-5]|[01]?\\d\\d?)\\.){3}(2[0-4]\\d|25[0-5]|[01]?\\d\\d?)$"
	return regexp.MustCompile(rule).MatchString(str)
}

//判断固定电话号码
//无视-符号匹配
func Tel(str string) bool {
	rule := "^(0\\d{2,3}(\\-)?)?\\d{7,8}$"
	return regexp.MustCompile(rule).MatchString(str)
}

//判断真实姓名
//只允许是汉字，并且长度为2-15位
func FullName(str string) bool {
	rule := `^[\p{Han}]{2,15}$`
	return regexp.MustCompile(rule).MatchString(str)
}
