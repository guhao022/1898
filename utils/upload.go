package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
)

func Upload(r *http.Request) (string, error) {

	fmt.Println("client:", r.RemoteAddr)

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	defer func() {
		if err := file.Close(); err != nil {
			println(err.Error())
			return
		}
	}()


	//文件名
	fileName := fileHeader.Filename
	if fileName == "" {
		return "", errors.New("Param filename cannot be null.")
	}
	//文件内容
	bytes, err := ioutil.ReadAll(file)

	//写到服务端本地文件中
	outputFilePath := "/var/www/html/1898/public/image/" + fileName
	err = ioutil.WriteFile(outputFilePath, bytes, os.ModePerm)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

func SHA1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

func is_valid_image(file *os.File) bool {
	return true
}
