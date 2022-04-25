package main

import (
	"fmt"
	"strings"
)

func after(value string, a string) string {
	// Get substring after a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}

func canConstruct(targetString string, stringList []string, memo map[string]bool) bool {
	if v, ok := memo[targetString]; ok {
		return v
	}
	if targetString == "" {
		return true
	}
	for _, s := range stringList {
		if strings.HasPrefix(targetString, s) {
			suffix := strings.TrimPrefix(targetString, s)
			if canConstruct(suffix, stringList, memo) == true {
				memo[targetString] = true
				return true
			}
		}
	}
	memo[targetString] = false
	return false
}

func main() {
	var memo = make(map[string]bool)
	//fmt.Println(canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))
	//fmt.Println(canConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))
	//fmt.Println(canConstruct("enterpotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))
	fmt.Println(canConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "g", "eeeeeeee", "n"}, memo))
}
