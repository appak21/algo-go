package main

import "fmt"

func fib(n int64) int64 {
	var f0 int64 = 0
	var f1 int64 = 1
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	rem := n % 60
	if rem == 0 {
		return 0
	}
	var i int64
	for i = 2; i < rem+3; i++ {
		f0, f1 = f1, (f0+f1)%60
	}
	return f1 - 1 //sum
}

func main() {
	var n int64
	fmt.Scan(&n)
	fmt.Println(fib(n) % 10)
}
