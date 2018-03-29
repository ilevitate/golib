package random

import (
	"math/rand"
	"time"

	"github.com/shopspring/decimal"
)

//指定范围内的随机数字
func RandInt(min int, max int) int {
	rand.Seed(rand.Int63n(time.Now().UnixNano()))
	return min + rand.Intn(max-min)
}

//指定范围的随机float值
func RandFloat(min, max float64, l int32) float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	f := r.Float64()*(max-min) + min
	result, _ := decimal.NewFromFloat(f).Truncate(l).Float64()
	return result
}
