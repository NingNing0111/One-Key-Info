package util

import (
	"log"
	"os"
)

func GetContent(htmlFile string) (string, error) {
	htmlByte, err := os.ReadFile(htmlFile)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return string(htmlByte), nil
}
