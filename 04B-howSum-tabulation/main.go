package main

import "fmt"

func howSum(targetSum int, numbers []int) []int {

	tab := make([][]int, targetSum+1)
	tab[0] = []int{}

	for i := 0; i <= targetSum; i++ {
		current := tab[i]
		if current != nil {
			for _, num := range numbers {
				if i+num <= targetSum {
					tab[i+num] = append(current, num)
				}
			}
		}
	}

	return tab[targetSum]
}

func main() {
	fmt.Printf("%#v\n", howSum(7, []int{2, 3}))       // [3, 2, 2]
	fmt.Printf("%#v\n", howSum(7, []int{5, 3, 4, 7})) // [4, 3]
	fmt.Printf("%#v\n", howSum(7, []int{2, 4}))       // null
	fmt.Printf("%#v\n", howSum(8, []int{2, 3, 5}))    // [2, 2, 2, 2]
	fmt.Printf("%#v\n", howSum(300, []int{7, 14}))    // null
}
