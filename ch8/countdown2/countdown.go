// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 244.

// Countdown implements the countdown for a rocket launch.
package main

import (
	"fmt"
	"os"
	"time"
)

//!+

func main() {
	// ...create abort channel...

	//!-

	//!+abort
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
	//!-abort

	//!+
	fmt.Println("Commencing countdown.  Press return to abort.")
	countdown := 0
	select {
	case <-time.After(10 * time.Second):
		// Do nothing.
		fmt.Println(countdown)
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	default:
		countdown++
		fmt.Println(countdown)
	}
	launch()
}

//!-

func launch() {
	fmt.Println("Lift off!")
}
