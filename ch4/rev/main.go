// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array

	//!+slice
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"

	b := [6]int{0, 1, 2, 3, 4, 5}
	reverse2(&b) //&s取s的指针
	//!-slice

	c := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(rotate(c, 2))

	str := []string{"a", "a", "b", "b", "d", "b", "c", "c"}
	fmt.Println(RemoveDuplicates(str))

	d := []byte("abc     a aaa     ccc  ddd d")
	e := RemoveSpace(d)
	fmt.Println(string(e))

	// Interactive test of reverse.
	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))
		}
		reverse(ints)
		fmt.Printf("%v\n", ints)
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!+rev
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//!-rev

func reverse2(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}

func rotate(s []int, r int) []int {
	lens := len(s)
	arr := make([]int, lens)
	for k := range s {
		index := r + k
		if index >= lens {
			index -= lens
		}
		arr[k] = s[index]
	}
	return arr
}

func RemoveDuplicates(str []string) []string {
	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			copy(str[i:], str[i+1:])
			str = str[:len(str)-1]
			i--
		}
	}
	return str
}

func RemoveSpace(s []byte) []byte {
	for i := 0; i < len(s)-1; i++ {
		if unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(s[i+1])) {
			// 结合remove函数
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
			i--
		}
	}
	return s
}
