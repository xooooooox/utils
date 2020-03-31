package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// ReadFile 读取文件内容, 返回字节切片
func ReadFile(filename string) (bs []byte, err error) {
	if !filepath.IsAbs(filename) {
		filename, err = filepath.Abs(filename)
		if err != nil {
			return
		}
	}
	_, err = os.Stat(filename)
	if err != nil {
		return
	}
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

// WriteFile 写入文件内容, 返回错误
func WriteFile(filename string, bs []byte) (err error) {
	if !filepath.IsAbs(filename) {
		filename, err = filepath.Abs(filename)
		if err != nil {
			return
		}
	}
	var file *os.File
	_, err = os.Stat(filename)
	if err != nil {
		file, err = os.Create(filename)
		if err != nil {
			return
		}
	} else {
		file, err = os.Open(filename)
		if err != nil {
			return
		}
	}
	defer file.Close()
	_, err = file.Write(bs)
	return
}
