// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma3(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

//!-

func comma2(s string) string {
	var buffer bytes.Buffer
	l := len(s)
	for i := 0; i < len(s); i++ {
		buffer.WriteString(string(s[i]))

		if (i+1)%3 == l%3 { //取余3可以得到第一个插入逗号的位置,后面依次+3即可
			buffer.WriteString(",")
		}
	}
	s = buffer.String()
	return s[:len(s)-1] // 末尾会多一个逗号,去掉
}

// 判断是否有正负号
// 判断是否有小数部分
func comma3(s string) string {
	var buffer bytes.Buffer

	// 获取正负号
	var symbol byte
	if s[0] == '-' || s[0] == '+' {
		symbol = s[0]
		s = s[1:]
	}

	// 将符号添加到返回的字符串中
	buffer.WriteByte(symbol)

	// 分离整数部分与小数部位
	arr := strings.Split(s, ".")
	s = arr[0]
	l := len(s)

	// 格式整数部分
	for i := 0; i < len(s); i++ {
		buffer.WriteString(string(s[i]))
		// 取余3可以得到第一个插入逗号的位置,后面依次+3即可,末尾不加","
		if (i+1)%3 == l%3 && i != l-1 {
			buffer.WriteString(",")
		}
	}

	// 存在小数部分
	if len(arr) > 1 {
		buffer.WriteString(".")
		buffer.WriteString(arr[1])
	}

	s = buffer.String()
	return s
}
