package main

import (
	"strings"
	"fmt"
	"bytes"
)

func main() {

	//func Join(a []string, sep string) string {
	strArray := []string{"hello", "world", "itcast"}

	strRes := strings.Join(strArray, "=")
	fmt.Printf("strRes : %s\n", strRes)


	//bytes.Jon
	//将二维切片使用一维切片链接起来，得到一个新的一维切片
	//func Join(s [][]byte, sep []byte) []byte {

	tmp:= [][]byte{
		[]byte("hello"),
		[]byte("world"),
		[]byte("itcast"),
	}

	bytesRes := bytes.Join(tmp, []byte{})
	fmt.Printf("bytesRes : %s\n", bytesRes)
}
