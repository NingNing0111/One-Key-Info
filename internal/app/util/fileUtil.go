package util

import (
	"log"
	"os"
)

// 获取指定路径的文件信息
func GetContent(htmlFile string) (string, error) {
	htmlByte, err := os.ReadFile(htmlFile)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return string(htmlByte), nil
}
