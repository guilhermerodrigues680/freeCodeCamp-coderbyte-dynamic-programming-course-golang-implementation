package main

import (
	"fmt"
	"strings"
)

func allConstructAux(target string, wordBank []string, memo map[string][][]string) [][]string {
	if res, found := memo[target]; found {
		return res
	}

	if len(target) == 0 {
		return [][]string{{}} // [[]]
	}

	allWays := [][]string{} // []
	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			suffix := target[len(word):]
			waysSuffix := allConstructAux(suffix, wordBank, memo)
			if len(waysSuffix) > 0 {
				for _, way := range waysSuffix {
					updatedWay := append([]string{word}, way...)
					allWays = append(allWays, updatedWay)
				}
			}
		}
	}

	memo[target] = allWays
	return memo[target]
}

func allConstruct(target string, wordBank []string) [][]string {
	memo := map[string][][]string{}
	return allConstructAux(target, wordBank, memo)
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
