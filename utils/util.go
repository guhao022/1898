package utils

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"

	"github.com/num5/storage"
)

// 闪存
func Flash(filename string, val interface{}) error  {
	Stor, err := storage.New("./",filename)

	if err != nil {
		return err
	}

	return Stor.Store(val)
}

// MD5加密
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))

	return hex.EncodeToString(h.Sum(nil))
}

// 匹配手机号
func MatchPhone (mobileNum string) bool {
	var reg = regexp.MustCompile("^((\\+86)|(86))?(1(([35][0-9])|[8][0-9]|[7][06789]|[4][579]))\\d{8}$")
	return reg.MatchString(mobileNum)
}
