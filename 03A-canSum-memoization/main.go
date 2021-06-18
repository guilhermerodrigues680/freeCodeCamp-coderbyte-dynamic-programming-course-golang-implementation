package main

import "fmt"

func canSumAux(targetSum int, numbers []int, memo map[int]bool) bool {
	if res, ok := memo[targetSum]; ok {
		return res
	}

	if targetSum < 0 {
		return false
	}

	if targetSum == 0 {
		return true
	}

	for _, num := range numbers {
		remainder := targetSum - num
		if canSumAux(remainder, numbers, memo) {
			memo[targetSum] = true
			return true
		}
	}

	memo[targetSum] = false
	return memo[targetSum]
}

func canSum(targetSum int, numbers []int) bool {
	memo := make(map[int]bool)
	return canSumAux(targetSum, numbers, memo)
}

func main() {
	fmt.Println(canSum(7, []int{2, 3}))       // true
	fmt.Println(canSum(7, []int{5, 3, 4, 7})) // true
	fmt.Println(canSum(7, []int{2, 4}))       // false
	fmt.Println(canSum(8, []int{2, 3, 5}))    // true
	fmt.Println(canSum(300, []int{7, 14}))    // false
}
