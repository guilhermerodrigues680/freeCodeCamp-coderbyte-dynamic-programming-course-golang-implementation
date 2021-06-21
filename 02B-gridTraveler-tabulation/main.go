package main

import (
	"fmt"
)

// func printTable(tab [][]uint64) {
// 	fmt.Println(strings.Repeat("-\t", len(tab[0])+2))
// 	for i := 0; i < len(tab); i++ {
// 		fmt.Print("|\t")
// 		for j := 0; j < len(tab[i]); j++ {
// 			fmt.Print(tab[i][j], "\t")
// 		}
// 		fmt.Println("|")
// 	}
// 	fmt.Println(strings.Repeat("-\t", len(tab[0])+2))
// }

func gridTraveler(m, n int) uint64 {
	tab := make([][]uint64, m+1)
	for i := 0; i < len(tab); i++ {
		tab[i] = make([]uint64, n+1)
	}

	tab[1][1] = 1

	// printTable(tab)

	// m = rows
	for i := 0; i <= m; i++ {
		// n = cols
		for j := 0; j <= n; j++ {
			current := tab[i][j]
			// move para baixo
			if i+1 <= m {
				tab[i+1][j] += current
			}
			// move direita
			if j+1 <= n {
				tab[i][j+1] += current
			}
		}
	}

	// printTable(tab)

	return tab[m][n]
}

func main() {
	fmt.Println(gridTraveler(1, 1))   // 1
	fmt.Println(gridTraveler(2, 3))   // 3
	fmt.Println(gridTraveler(3, 2))   // 3
	fmt.Println(gridTraveler(3, 3))   // 6
	fmt.Println(gridTraveler(18, 18)) // 2333606220
}
