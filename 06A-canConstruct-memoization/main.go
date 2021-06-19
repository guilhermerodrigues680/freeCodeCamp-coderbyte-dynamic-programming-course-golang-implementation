package main

import (
	"fmt"
	"strings"
)

func canConstructAux(target string, wordBank []string, memo map[string]bool) bool {
	if res, found := memo[target]; found {
		return res
	}

	if len(target) == 0 {
		return true
	}

	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			suffix := target[len(word):]
			if canConstructAux(suffix, wordBank, memo) {
				memo[target] = true
				return memo[target]
			}
		}
	}

	memo[target] = false
	return memo[target]
}

func canConstruct(target string, wordBank []string) bool {
	memo := make(map[string]bool)
	return canConstructAux(target, wordBank, memo)
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
