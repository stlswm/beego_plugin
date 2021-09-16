package app

import (
	"github.com/beego/beego/v2/server/web"
	"os"
	"strings"
)

// RuntimePath 获取运行时临时地址
func RuntimePath() string {
	path, _ := web.AppConfig.String("runtime.path")
	path = strings.TrimRight(strings.Replace(path, "\\", "/", -1), "/")
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, 0x0755)
		panic(err)
	}
	return path
}

// RuntimeSubPath 获取运行时临时地址子目录
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
