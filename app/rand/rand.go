package rand

import (
	"math/rand"
	"strings"
	"time"
)

const (
	KIND_NUM   = 0 // 纯数字
	KIND_LOWER = 1 // 小写字母
	KIND_UPPER = 2 // 大写字母
	KIND_ALL   = 3 // 数字、大小写字母
)

var numSeek string
var lowerSeek string
var upperSeek string
var allSeek string

func init() {
	numSeek = "0123456789"
	lowerSeek = "abcdefghijklmnopqrstuvwxyz"
	upperSeek = strings.ToUpper(lowerSeek)
	allSeek = numSeek + lowerSeek + upperSeek
}
func RangeNum(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randNum := r.Intn(max-min) + min
	return randNum
}

// 随机字符串
func RandomStr(l int) string {
	var randStr string
	seek := "0123456789"
	seek += "abcdefghijklmnopqrstuvwxyz"
	seek += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	totalLen := len(seek)
	for i := 0; i < l; i++ {
		k := RangeNum(0, totalLen)
		randStr += seek[k:(k + 1)]
	}
	return randStr
}
