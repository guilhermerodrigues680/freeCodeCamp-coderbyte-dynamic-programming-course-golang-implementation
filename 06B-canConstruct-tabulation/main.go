package main

import (
	"fmt"
)

func canConstruct(target string, wordBank []string) bool {

	tab := make([]bool, len(target)+1)
	tab[0] = true

	for i := 0; i <= len(target); i++ {
		current := tab[i]
		if current == true {
			for _, word := range wordBank {
				if i+len(word) <= len(target) {
					sub := target[i : i+len(word)]
					if sub == word {
						tab[i+len(word)] = true
					}
				}
			}
		}
	}

	return tab[len(target)]
}

func main() {
	fmt.Printf("%#v\n", canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))                  // true
	fmt.Printf("%#v\n", canConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))   // false
	fmt.Printf("%#v\n", canConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"})) // true
	fmt.Printf("%#v\n", canConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{
		"e",
		"ee",
		"eee",
		"eeee",
		"eeeee",
		"eeeeee",
	})) // false
}
