package http

import (
	"fmt"
	"testing"
)

func Test_getKeyInfo(t *testing.T) {
	keyInfo := GetKeyInfo("sk-******", "https://api.mnzdna.xyz/")
	fmt.Println(keyInfo)
}
