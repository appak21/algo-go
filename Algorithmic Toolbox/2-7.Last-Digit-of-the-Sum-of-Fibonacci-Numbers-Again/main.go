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
		f0, f1 = f1, (f0+f1)%10
	}
	return f1 - 1 //sum
}

func main() {
	var m, n int64
	fmt.Scan(&m, &n)
	final := abs(fib(n) - fib(m-1) + 10)
	fmt.Println(final % 10)
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
