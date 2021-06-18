package main

import (
	"fmt"
	"math"
)

func bestSum(targetSum int, numbers []int) []int {
	if targetSum < 0 {
		return nil
	}

	if targetSum == 0 {
		return make([]int, 0)
	}

	possibleRes := make(map[int][][]int)
	for _, num := range numbers {
		bs := bestSum(targetSum-num, numbers)
		if bs != nil {
			b := append(bs, num)
			length := len(b)
			_, exists := possibleRes[length]
			if !exists {
				possibleRes[length] = make([][]int, 0)
			}

			possibleRes[length] = append(possibleRes[length], b)
		}
	}

	if len(possibleRes) == 1 {
		for _, v := range possibleRes {
			return v[0]
		}
	}

	minKey := math.MaxInt32
	if len(possibleRes) > 1 {
		for k := range possibleRes {
			if k < minKey {
				minKey = k
			}
		}
		return possibleRes[minKey][0]
	}

	return nil
}

func main() {
	fmt.Printf("%#v\n", bestSum(7, []int{2, 3}))       // [3, 2, 2]
	fmt.Printf("%#v\n", bestSum(7, []int{5, 3, 4, 7})) // [4, 3]
	fmt.Printf("%#v\n", bestSum(7, []int{2, 4}))       // null
	fmt.Printf("%#v\n", bestSum(8, []int{2, 3, 5}))    // [2, 2, 2, 2]
	fmt.Printf("%#v\n", bestSum(300, []int{7, 14}))    // null
}
