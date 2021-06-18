package main

import "fmt"

func howSum(targetSum int, numbers []int) []int {
	if targetSum < 0 {
		return nil
	}

	if targetSum == 0 {
		emptyArr := make([]int, 0)
		return emptyArr
	}

	for _, num := range numbers {
		remainder := targetSum - num
		res := howSum(remainder, numbers)
		if res != nil {
			res := append(res, num)
			return res
		}
	}

	return nil
}

func main() {
	fmt.Printf("%#v\n", howSum(7, []int{2, 3}))       // [3, 2, 2]
	fmt.Printf("%#v\n", howSum(7, []int{5, 3, 4, 7})) // [4, 3]
	fmt.Printf("%#v\n", howSum(7, []int{2, 4}))       // null
	fmt.Printf("%#v\n", howSum(8, []int{2, 3, 5}))    // [2, 2, 2, 2]
	fmt.Printf("%#v\n", howSum(300, []int{7, 14}))    // null
}
