package aes128

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/stlswm/goencrypt/aes128"
)

// Encrypt 加密
func Encrypt(src string) (error, string) {
	key, _ := web.AppConfig.String("encrypt.key")
	iv, _ := web.AppConfig.String("encrypt.iv")
	return aes128.Encrypt(src, key, iv)
}

// Decrypt 解密
func Decrypt(src string) (error, string) {
	key, _ := web.AppConfig.String("encrypt.key")
	iv, _ := web.AppConfig.String("encrypt.iv")
	return aes128.Decrypt(src, key, iv)
}
