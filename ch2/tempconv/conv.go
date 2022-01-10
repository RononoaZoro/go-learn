// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package tempconv

import (
	"fmt"
)

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CTok(c Celsius) Kelvins { return Kelvins(c - AbsoluteZeroC) }
func KtoC(k Kelvins) Celsius { return Celsius(k + Kelvins(AbsoluteZeroC)) }

//!-

// func main() {
// 	c := AbsoluteZeroC
// 	fmt.Print(CToF(c).String())
// }
