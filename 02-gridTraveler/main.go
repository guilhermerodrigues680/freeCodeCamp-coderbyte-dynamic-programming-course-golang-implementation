package main

import "fmt"

func gridTraveler(m, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	if m == 1 && n == 1 {
		return 1
	}

	// total para baixo + total para direita
	return gridTraveler(m-1, n) + gridTraveler(m, n-1)
}

func main() {
	fmt.Println(gridTraveler(1, 1))   // 1
	fmt.Println(gridTraveler(2, 3))   // 3
	fmt.Println(gridTraveler(3, 2))   // 3
	fmt.Println(gridTraveler(3, 3))   // 6
	fmt.Println(gridTraveler(18, 18)) // 2333606220
}
