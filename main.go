package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var whatIsIt string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)
	str := string(sd)
	sLen := len(str)
	result := make([]byte, sLen)
	for i := range str {
		result[i] = sd[sLen-1-i]
	}
	whatIsIt = string(result)
	fmt.Println(whatIsIt)
}
