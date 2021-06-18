package main

import "fmt"

func howSumAux(targetSum int, numbers []int, memo map[int][]int) []int {
	if res, ok := memo[targetSum]; ok {
		return res
	}

	if targetSum < 0 {
		return nil
	}

	if targetSum == 0 {
		emptyArr := make([]int, 0)
		return emptyArr
	}

	for _, num := range numbers {
		remainder := targetSum - num
		res := howSumAux(remainder, numbers, memo)
		if res != nil {
			memo[targetSum] = append(res, num)
			return memo[targetSum]
		}
	}

	memo[targetSum] = nil
	return memo[targetSum]
}

func howSum(targetSum int, numbers []int) []int {
	memo := make(map[int][]int)
	return howSumAux(targetSum, numbers, memo)
}

func main() {
	fmt.Printf("%#v\n", howSum(7, []int{2, 3}))       // [3, 2, 2]
	fmt.Printf("%#v\n", howSum(7, []int{5, 3, 4, 7})) // [4, 3]
	fmt.Printf("%#v\n", howSum(7, []int{2, 4}))       // null
	fmt.Printf("%#v\n", howSum(8, []int{2, 3, 5}))    // [2, 2, 2, 2]
	fmt.Printf("%#v\n", howSum(300, []int{7, 14}))    // null
}
