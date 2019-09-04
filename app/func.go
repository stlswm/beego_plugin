package app

import (
	"github.com/astaxie/beego"
	"math/rand"
	"os"
	"strings"
	"time"
)

// 获取运行时临时地址
func RuntimePath() string {
	path := beego.AppConfig.String("runtime.path")
	path = strings.TrimRight(strings.Replace(path, "\\", "/", -1), "/")
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, 0x0755)
		panic(err)
	}
	return path
}

// 获取运行时临时地址子目录
func RuntimeSubPath(path string) string {
	path = strings.TrimLeft(strings.Replace(path, "\\", "/", -1), "/")
	tmpPath := RuntimePath()
	path = tmpPath + "/" + path
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, 0x0755)
		panic(err)
	}
	return path
}

// 指定区间的随机数
func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}

// 随机字符串
func RandStr(l int) string {
	var randStr string
	seek := "0123456789"
	seek += "abcdefghijklmnopqrstuvwxyz"
	seek += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	totalLen := len(seek)
	for i := 0; i < l; i++ {
		k := GenerateRangeNum(0, totalLen-1)
		randStr += seek[k:(k + 1)]
	}
	return randStr
}
