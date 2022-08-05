package main

import "fmt"

func pisano(m int) int {
	prev, curr, res := 0, 1, 0
	for i := 0; i < m*m; i++ {
		prev, curr = curr, (prev+curr)%m
		if prev == 0 && curr == 1 {
			res = i + 1
		}
	}
	return res
}

func fibonacciModulo(n int64, m int) int {
	pisano := pisano(m)
	n = n % int64(pisano)
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	prev, curr := 0, 1
	for i := 0; i < int(n)-1; i++ {
		prev, curr = curr, (prev+curr)%m
	}
	return curr % m
}

func main() {
	var n int64
	var m int
	fmt.Scan(&n, &m)
	fmt.Println(fibonacciModulo(n, m))
}
