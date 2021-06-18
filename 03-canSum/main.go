package main

import "fmt"

func canSum(targetSum int, numbers []int) bool {
	if targetSum < 0 {
		return false
	}

	if targetSum == 0 {
		return true
	}

	for _, num := range numbers {
		remainder := targetSum - num
		if canSum(remainder, numbers) {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(canSum(7, []int{2, 3}))       // true
	fmt.Println(canSum(7, []int{5, 3, 4, 7})) // true
	fmt.Println(canSum(7, []int{2, 4}))       // false
	fmt.Println(canSum(8, []int{2, 3, 5}))    // true
	fmt.Println(canSum(300, []int{7, 14}))    // false
}
