/*
 SYNOPSIS:

 With GCC >= 4.6:
 # gccgo gcd.go -o gcd-go
 # ./gcd-go 11 22 33 44 121

 With Google Go (http://golang.org/):
 # go run gcd.go 11 22 33 44 121
 # or, if you want to play with the binary
 # go build -o gcd-go gcd.go
 # ./gcd-go 11 22 33 44 121

 GCC makes dynamically linked binary,
 but Google Go - statically linked

*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func gcd2(a, b uint64) uint64 {
	if b == 0 {
		return a
	}
	return gcd2(b, a%b)
}

func gcdn(ns []uint64) (r uint64) {
	for _, v := range ns {
		r = gcd2(r, v)
	}
	return
}

func main() {
	if len(os.Args) == 1 {
		return
	}
	var ns []uint64 // We have garbage collector!
	for _, arg := range os.Args[1:] {
		// Drop the second return value (error code):
		v, _ := strconv.ParseUint(arg, 0, 64)
		ns = append(ns, v)
	}
	fmt.Printf("%v\n", gcdn(ns))
}
