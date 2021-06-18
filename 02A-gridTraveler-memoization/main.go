package main

import (
	"fmt"
	"strconv"
)

func gridTravelerAux(m int, n int, memo map[string]int) int {
	keyMN := strconv.Itoa(m) + "," + strconv.Itoa(n)
	if res, ok := memo[keyMN]; ok {
		return res
	}

	if m <= 0 || n <= 0 {
		return 0
	}

	if m == 1 && n == 1 {
		return 1
	}

	memo[keyMN] = gridTravelerAux(m-1, n, memo) + gridTravelerAux(m, n-1, memo)
	return memo[keyMN]
}

func gridTraveler(m, n int) int {
	memo := make(map[string]int)
	return gridTravelerAux(m, n, memo)
}

func main() {
	fmt.Println(gridTraveler(1, 1))   // 1
	fmt.Println(gridTraveler(2, 3))   // 3
	fmt.Println(gridTraveler(3, 2))   // 3
	fmt.Println(gridTraveler(3, 3))   // 6
	fmt.Println(gridTraveler(18, 18)) // 2333606220
}
