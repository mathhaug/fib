package main

import (
	"fmt"

	"github.com/mathhaug/fib/fibo"
)

func main() {
	n := 10
	result := fibo.Fib(n)
	fmt.Printf("Fib(%d) = %d\n", n, result)
}
