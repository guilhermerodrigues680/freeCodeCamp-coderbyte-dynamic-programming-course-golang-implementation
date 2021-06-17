package main

import "fmt"

func fibAux(n int, memo map[int]int64) int64 {
	if memoN, ok := memo[n]; ok {
		return memoN
	}

	if n <= 2 {
		return 1
	}

	memo[n] = fibAux(n-1, memo) + fibAux(n-2, memo)
	return memo[n]
}

func fib(n int) int64 {
	memo := make(map[int]int64)
	return fibAux(n, memo)
}

func main() {
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(9))
	fmt.Println(fib(50))
	fmt.Println(fib(32767))
}
