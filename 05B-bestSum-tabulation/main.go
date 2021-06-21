package main

import (
	"fmt"
)

func bestSum(targetSum int, numbers []int) []int {
	tab := make([][]int, targetSum+1)
	tab[0] = []int{}

	for i := 0; i <= targetSum; i++ {
		current := tab[i]
		if current != nil {
			for _, num := range numbers {
				if i+num <= targetSum {
					valINum := tab[i+num]
					if valINum == nil {
						tab[i+num] = append(current, num)
					} else {
						updatedVal := append(current, num)
						if len(updatedVal) < len(valINum) {
							tab[i+num] = updatedVal
						}
					}
				}
			}
		}
	}

	return tab[targetSum]
}

func main() {
	fmt.Printf("%#v\n", bestSum(7, []int{5, 3, 4, 7}))    // [7]
	fmt.Printf("%#v\n", bestSum(8, []int{2, 3, 5}))       // [3, 5]
	fmt.Printf("%#v\n", bestSum(8, []int{1, 4, 5}))       // [4, 4]
	fmt.Printf("%#v\n", bestSum(100, []int{1, 2, 5, 25})) // [25, 25, 25, 25]
}
