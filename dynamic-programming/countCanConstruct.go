package main

import (
	"fmt"
	"strings"
)

func countCanConstruct(targetString string, stringList []string, memo map[string]int) int {
	if v, ok := memo[targetString]; ok {
		return v
	}
	if targetString == "" {
		return 1
	}
	var count = 0
	for _, s := range stringList {
		if strings.HasPrefix(targetString, s) {
			suffix := strings.TrimPrefix(targetString, s)
			count += countCanConstruct(suffix, stringList, memo)
		}
	}
	memo[targetString] += count
	return count
}

func main() {
	var memo = make(map[string]int)
	//fmt.Println(countCanConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}, memo))
	//fmt.Println(countCanConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}, memo))
	//fmt.Println(countCanConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, memo))
	fmt.Println(countCanConstruct("enterpotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}, memo))
	//fmt.Println(countCanConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "g", "eeeeeeee", "f"}, memo))
}
