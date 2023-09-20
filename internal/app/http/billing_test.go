package http

import (
	"fmt"
	"testing"
)

func Test_getKeyInfo(t *testing.T) {
	keyInfo := GetKeyInfo("sk-***y", "https://api.mnzdna.xyz/")
	fmt.Println(keyInfo)
}
