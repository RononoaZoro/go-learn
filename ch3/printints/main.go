// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 74.

// Printints demonstrates the use of bytes.Buffer to format a string.
// x << n 等价于 x * 2^n
package main

import (
	"bytes"
	"fmt"
)

//!+
// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}

//!-

func isReverse(a, b string) bool {
	// 长度不一样直接返回false
	if len(a) != len(b) {
		return false
	}
	// 用于记录每个字符串出现的次数
	m := make(map[rune]int)
	n := make(map[rune]int)
	// 以字符串Unicode码作为map的Key
	for _, v := range a {
		m[v]++
	}
	for _, v := range b {
		n[v]++
	}
	// 判断相同下标值是否相同
	for i, v := range m {
		if n[i] != v {
			return false
		}
	}
	return true
}
