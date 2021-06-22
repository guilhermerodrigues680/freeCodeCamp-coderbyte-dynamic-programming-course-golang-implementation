package main

import (
	"fmt"
)

func allConstruct(target string, wordBank []string) [][]string {
	tab := make([][][]string, len(target)+1)
	tab[0] = [][]string{{}} // [[]]
	for i := 1; i < len(tab); i++ {
		tab[i] = [][]string{} // []
	}

	for i := 0; i <= len(target); i++ {
		current := tab[i]
		for _, word := range wordBank {
			if i+len(word) <= len(target) {
				sub := target[i : i+len(word)]
				if sub == word {
					for _, v := range current {
						newCombination := append(v, word)
						tab[i+len(word)] = append(tab[i+len(word)], newCombination)
					}
				}
			}
		}
	}

	return tab[len(target)]
}

func main() {
	fmt.Println(allConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))
	// [
	// 	[purp le]
	// 	[p ur p le]
	// ]
	fmt.Println(allConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}))
	// [
	// 	[ab cd ef]
	// 	[ab c def]
	// 	[abc def]
	// 	[abcd ef]
	// ]
	fmt.Println(allConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))
	// []
	fmt.Println(allConstruct("aaaaaaaaaaaaaaaaaaaaaaaaaaz", []string{"a", "aa", "aaa", "aaaa", "aaaaa"}))
	// []
}
