package app

import (
	"github.com/astaxie/beego"
	"os"
	"strings"
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
