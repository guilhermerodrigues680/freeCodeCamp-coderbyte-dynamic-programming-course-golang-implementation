package main

import "fmt"

func fib(n int) uint64 {
	tab := make([]uint64, n+1)
	tab[1] = 1

	// Minha implementação
	// for i := 2; i <= n; i++ {
	// 	tab[i] = tab[i-1] + tab[i-2]
	// }

	// Implementação usada no curso (lança `panic: runtime error: index out of range`)
	// for i := 0; i <= n; i++ {
	// 	tab[i+1] += tab[i]
	// 	tab[i+2] += tab[i]
	// }

	// Implementação usada no curso com checagem de index
	for i := 0; i <= n; i++ {
		if i+1 <= n {
			tab[i+1] += tab[i]
		}

		if i+2 <= n {
			tab[i+2] += tab[i]
		}
	}

	return tab[n]
}

func main() {
	fmt.Println(fib(6))  // 8
	fmt.Println(fib(7))  // 13
	fmt.Println(fib(8))  // 21
	fmt.Println(fib(50)) // 12586269025
}
