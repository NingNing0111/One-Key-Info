package http

import (
	"fmt"
	"testing"
)

func Test_getKeyInfo(t *testing.T) {
	keyInfo := GetKeyInfo("sk-nkASlWIspEgYKSQK29FfB690D36b4937Ac30AeDf1b44F18e", "https://api.mnzdna.xyz/")
	fmt.Println(keyInfo)
}
