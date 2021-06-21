package main

import "fmt"

func canSum(targetSum int, numbers []int) bool {

	tab := make([]bool, targetSum+1)
	tab[0] = true

	for i := 0; i <= targetSum; i++ {
		current := tab[i]
		if current == true {
			for _, num := range numbers {
				if i+num <= targetSum {
					tab[i+num] = true
				}
			}
		}
	}

	return tab[targetSum]
}

func main() {
	fmt.Println(canSum(7, []int{2, 3}))       // true
	fmt.Println(canSum(7, []int{5, 3, 4, 7})) // true
	fmt.Println(canSum(7, []int{2, 4}))       // false
	fmt.Println(canSum(8, []int{2, 3, 5}))    // true
	fmt.Println(canSum(300, []int{7, 14}))    // false
}
