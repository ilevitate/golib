package maths

import "math/rand"

//指定值在指定0至rangeValue区间随机命中的结果
func Probability(value int, rangeValue int) bool {
	//获取一个0-100的随机值
	r := rand.Intn(rangeValue)
	//如果随机在几率范围内
	if r <= value {
		return true
	}
	return false
}