// @Author: su
// @Date: 2021/6/22 15:36

package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/md4"
)

func Md4Encode(src string) string {
	src_byte := []byte(src)
	md4_new := md4.New()
	md4Bytes := md4_new.Sum(src_byte)
	return hex.EncodeToString(md4Bytes)
}

func Md5Encode(src string) string {
	src_byte := []byte(src)
	// md5计算
	md5_new := md5.New()
	md5Bytes := md5_new.Sum(src_byte)
	return hex.EncodeToString(md5Bytes)
}

func Sha256Encoding(src string) string {
	sha256bytes := sha256.Sum256([]byte(src))
	return hex.EncodeToString(sha256bytes[:])
}

func main() {
	// 生成md4
	src := "hello world"
	md4_src := Md4Encode(src)
	fmt.Println("Md4Encode:", md4_src)

	// 生成md5
	md5_src := Md5Encode(src)
	fmt.Println("Md5Encode:", md5_src)

	// 生成sha256
	sha256_src := Sha256Encoding(src)
	fmt.Println("sha256:", sha256_src)
}
