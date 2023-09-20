package util

import (
	"encoding/json"
	"log"
)

// Parse 字符串转任意结构体
func Parse(jsonStr string, v any) {
	err := json.Unmarshal([]byte(jsonStr), v)
	if err != nil {
		log.Println("JSON解析失败：", err)
		log.Fatalln("字段为：", jsonStr)
	}
}
