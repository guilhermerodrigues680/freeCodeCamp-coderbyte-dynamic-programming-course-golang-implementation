package main

import (
	"fmt"
)

func countConstruct(target string, wordBank []string) int {
	tab := make([]int, len(target)+1)
	tab[0] = 1

	for i := 0; i <= len(target); i++ {
		current := tab[i]
		for _, word := range wordBank {
			if i+len(word) <= len(target) {
				sub := target[i : i+len(word)]
				if sub == word {
					tab[i+len(word)] += current
				}
			}
		}
	}

	return tab[len(target)]
}

func main() {
	fmt.Printf("%#v\n", countConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))                  // 2
	fmt.Printf("%#v\n", countConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))                  // 1
	fmt.Printf("%#v\n", countConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))   // 0
	fmt.Printf("%#v\n", countConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"})) // 4
	fmt.Printf("%#v\n", countConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{
		"e",
		"ee",
		"eee",
		"eeee",
		"eeeee",
		"eeeeee",
	})) // 0
}
