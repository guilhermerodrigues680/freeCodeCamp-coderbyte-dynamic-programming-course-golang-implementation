package main

import (
	"fmt"
	"strings"
)

func canConstruct(target string, wordBank []string) bool {
	if len(target) == 0 {
		return true
	}

	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			suffix := target[len(word):]
			if canConstruct(suffix, wordBank) {
				return true
			}
		}
	}

	return false
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
