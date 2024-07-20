package main

import (
	"fmt"
)

var g int

func main() {
	divisors(4)

}
func divisors(n int) int {
	for i := 1; i <= n; i++ {

		if n%i == 0 {
			g += 1
		}
	}
	fmt.Println(g)
	return g

}
