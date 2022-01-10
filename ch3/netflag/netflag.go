// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 77.

// Netflag demonstrates an integer type used as a bit field.
package main

import (
	"fmt"
	. "net"
)

//!+
func IsUp(v Flags) bool { return v&FlagUp == FlagUp }
func TurnDown(v *Flags) {
	*v &^= FlagUp
	//等同于a&(^b)
}
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b\n", FlagUp)
	fmt.Printf("%b\n", FlagBroadcast)
	fmt.Printf("%b\n", FlagLoopback)
	fmt.Printf("%b\n", FlagPointToPoint)
	fmt.Printf("%b\n", FlagMulticast)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}

const (
	KB = 1000
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776 (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424 (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

// const (
// 	KB = 1000
// 	MB = MiB - MiB%(KB*KB)
// 	GB = GiB - GiB%(MB*KB)
// 	TB = TiB - TiB%(GB*KB)
// 	PB = PiB - PiB%(TB*KB)
// 	EB = EiB - EiB%(PB*KB)
// 	ZB = ZiB - ZiB%(EB*KB)
// 	YB = YiB - YiB%(ZB*KB)
// )

//!-
