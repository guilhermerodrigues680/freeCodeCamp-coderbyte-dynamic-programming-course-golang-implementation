package main

import (
	"fmt"
)

func bestSum(targetSum int, numbers []int) []int {
	if targetSum < 0 {
		return nil
	}

	if targetSum == 0 {
		return make([]int, 0)
	}

	shortestCombination := []int(nil)
	for _, num := range numbers {
		remainder := targetSum - num
		remainderCombination := bestSum(remainder, numbers)
		if remainderCombination != nil {
			combination := append(remainderCombination, num)
			if shortestCombination == nil || len(combination) < len(shortestCombination) {
				shortestCombination = combination
			}
		}
	}

	return shortestCombination
}

func main() {
	fmt.Printf("%#v\n", bestSum(7, []int{5, 3, 4, 7}))    // [7]
	fmt.Printf("%#v\n", bestSum(8, []int{2, 3, 5}))       // [3, 5]
	fmt.Printf("%#v\n", bestSum(8, []int{1, 4, 5}))       // [4, 4]
	fmt.Printf("%#v\n", bestSum(100, []int{1, 2, 5, 25})) // [25, 25, 25, 25]
}
