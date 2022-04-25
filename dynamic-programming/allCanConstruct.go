package main

import (
	"fmt"
	"strings"
)

func allCanConstruct(targetString string, stringList []string, memo map[string][][]string) [][]string {
	if v, ok := memo[targetString]; ok {
		return v
	}
	if targetString == "" {
		return [][]string{}
	}
	var combinations [][]string
	for _, s := range stringList {
		if strings.HasPrefix(targetString, s) {
			suffix := strings.TrimPrefix(targetString, s)
			combination := allCanConstruct(suffix, stringList, memo)
			if combination != nil {
				var newCombinationList [][]string
				if len(combination) == 0 {
					var newList []string
					newList = append(newList, s)
					newCombinationList = append(newCombinationList, newList)
				} else {
					for _, com := range combination {
						var newList []string
						newList = append(newList, s)
						newList = append(newList, com...)
						newCombinationList = append(newCombinationList, newList)
					}
				}
				combinations = append(combinations, newCombinationList...)
			}
		}

	}
	memo[targetString] = combinations
	return combinations
}

func main() {
	var memo = make(map[string][][]string)
	//fmt.Println(allCanConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}, memo))
	//fmt.Println(allCanConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}))
	//fmt.Println(allCanConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))
	//fmt.Println(allCanConstruct("enterpotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}, memo))
	fmt.Println(allCanConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", []string{"e", "ee", "eee", "eeee", "eeeee", "g", "eeeeeeee", "f"}, memo))
}
