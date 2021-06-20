package main

import (
	"fmt"
	"strings"
)

func countConstructAux(target string, wordBank []string, memo map[string]int) int {
	if res, found := memo[target]; found {
		return res
	}

	if len(target) == 0 {
		return 1
	}

	count := 0
	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			suffix := target[len(word):]
			countSuffix := countConstructAux(suffix, wordBank, memo)
			count += countSuffix
		}
	}

	memo[target] = count
	return memo[target]
}

func countConstruct(target string, wordBank []string) int {
	memo := make(map[string]int)
	return countConstructAux(target, wordBank, memo)
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
