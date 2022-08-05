package main

import "fmt"

func fib(n int64) int {
	if n == 0 {
		return 0
	}
	f := make([]int, 60)
	last := make([]int, 60)
	f[0], f[1] = 0, 1
	last[0], last[1] = 0, 1
	for i := 2; i < 60; i++ {
		f[i] = f[i-1] + f[i-2]
		last[i] = f[i] % 10
	}
	rem := int(n % 60)
	s := last[rem]*last[rem] + last[rem]*last[int((n-1)%60)]
	return s % 10
}

func main() {
	var n int64
	fmt.Scan(&n)
	fmt.Println(fib(n))
}
