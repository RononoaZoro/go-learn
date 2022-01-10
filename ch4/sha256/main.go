// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"crypto/sha256" //!+
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8

	fmt.Println(compareSha256("X", "x"))
}

//找出两个字符串的不同bit位数量
func compareSha256(str1 string, str2 string) int {
	a := sha256.Sum256([]byte(str1))
	b := sha256.Sum256([]byte(str2))
	num := 0
	// 循环字节数组
	for i := 0; i < len(a); i++ {
		for m := 1; m <= 8; m++ { // 对比字节是否相同
			if (a[i] >> uint(m)) != (b[i] >> uint(m)) {
				num++
			}
		}
	}
	return num
}

//!-
