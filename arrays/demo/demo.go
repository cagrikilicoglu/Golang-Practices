package demo

import (
	"strings"
)


func UpdateValue(arr *[3]string) [3]string{
	for i,v := range arr {
		arr[i] = strings.Join([]string{v, "test"}," ")
	}
	return *arr
}