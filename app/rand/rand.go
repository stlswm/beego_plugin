package rand

import (
	"math/rand"
	"time"
)

const (
	KindNum    = 0 // 纯数字
	KindLower  = 1 // 小写字母
	KindUpper  = 2 // 大写字母
	KindLetter = 3 // 所有字母
	KindAll    = 4 // 数字、大小写字母
)

func RangeNum(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randNum := r.Intn(max-min) + min
	return randNum
}

// 随机字符串
func RandomStr(kind, l int) string {
	var randStr string
	var seek string
	switch kind {
	case KindNum:
		seek = "0123456789"
	case KindLower:
		seek = "abcdefghijklmnopqrstuvwxyz"
	case KindUpper:
		seek = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case KindLetter:
		seek = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	default:
		seek = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	totalLen := len(seek)
	for i := 0; i < l; i++ {
		k := RangeNum(0, totalLen)
		randStr += seek[k:(k + 1)]
	}
	return randStr
}
