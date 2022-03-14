package tdd 

import "strings"
var vowels = map[string]bool{
	"a" : true,
	"e" : true,
	"i" : true,
	"o" : true,
	"u" : true,
}

func VowelCount(str string) uint {
	var count uint
	str = strings.ToLower(str)
	for _, char := range str {
		if vowels[string(char)]{
			count++
		}
	}
	return count
}