package aes128

import (
	"github.com/astaxie/beego"
	"github.com/stlswm/goencrypt/aes128"
)

// 加密
func Encrypt(src string) string {
	key := beego.AppConfig.String("encrypt.key")
	iv := beego.AppConfig.String("encrypt.iv")
	return aes128.Encrypt(src, key, iv)
}

// 解密
func Decrypt(src string) string {
	key := beego.AppConfig.String("encrypt.key")
	iv := beego.AppConfig.String("encrypt.iv")
	return aes128.Decrypt(src, key, iv)
}
