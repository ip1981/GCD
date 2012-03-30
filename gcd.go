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

// Both Google Go and GCC issue an error "imported and not used",
// if imported and not used :-)
import (
	"flag"
	"fmt"
	"strconv"
)

func gcd2(a, b uint64) uint64 {
	if b == 0 {
		return a
	}
	/* 6g issues an error "function ends without a return statement",
	   if we use if ... {... return} else {... return}.
	   But GCC doesn't.
	*/
	return gcd2(b, a%b)
}

func gcdn(ns []uint64) (r uint64) {
	for i := range ns {
		r = gcd2(r, ns[i])
	}
	return
}

func main() {
	flag.Parse() // without this 6g will give flag.NArg() = 0 next (WTF?)
	n := flag.NArg()
	if n > 0 {
		ns := make([]uint64, n) // We have garbage collector!

		// Or: for i := range ns, since range of ns is equal to flag.NArg()
		for i := 0; i < n; i++ {
			// Drop the second return value (error code):
			ns[i], _ = strconv.ParseUint(flag.Arg(i), 0, 64)
		}

		g := gcdn(ns)
		fmt.Printf("%v\n", g)
	}
}