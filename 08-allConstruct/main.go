package main

import (
	"fmt"
	"strings"
)

func allConstruct(target string, wordBank []string) [][]string {
	if len(target) == 0 {
		return [][]string{{}} // [[]]
	}

	allWays := [][]string{} // []
	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			suffix := target[len(word):]
			waysSuffix := allConstruct(suffix, wordBank)
			if len(waysSuffix) > 0 {
				for _, way := range waysSuffix {
					updatedWay := append([]string{word}, way...)
					allWays = append(allWays, updatedWay)
				}
			}
		}
	}

	return allWays
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
