// @Author: su
// @Date: 2021/6/18 16:13

package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// 字符串加密
	input := []byte("hello world")
	eStr := base64.StdEncoding.EncodeToString(input)
	fmt.Println("字符串加密:", eStr)

	// 字符串解密
	decodeString, err := base64.StdEncoding.DecodeString(eStr)
	if err != nil {
		fmt.Println("decoding error")
		return
	}
	fmt.Println("字符串解密:", string(decodeString))

	// URL加密
	url_input := []byte("https://www.baidu.com")
	u_str:= base64.URLEncoding.EncodeToString(url_input)
	fmt.Println("url加密:", u_str)

	// URL解密
	decodeUrl, err := base64.URLEncoding.DecodeString(u_str)
	if err != nil {
		fmt.Println("url decoding error")
		return
	}
	fmt.Println("URL解密:", string(decodeUrl))
}
