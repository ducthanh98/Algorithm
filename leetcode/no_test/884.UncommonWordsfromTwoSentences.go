package main

import (
	"fmt"
	"strings"
)

func uncommonFromSentences(s1 string, s2 string) []string {
	arr1 := strings.Split(s1," ")
	arr2 := strings.Split(s2," ")
	mapArr1 := make(map[string]int)
	results := make([]string,0)
	for i := 0; i < len(arr1); i++ {
		mapArr1[arr1[i]]++
	}

	for i := 0; i < len(arr2); i++ {
		mapArr1[arr2[i]]++
	}

	for i := 0; i < len(arr1); i++ {
		if mapArr1[arr1[i]] == 1 {
			results = append(results,arr1[i])
		}
	}

	for i := 0; i < len(arr2); i++ {
		if mapArr1[arr2[i]] == 1{
			results = append(results,arr2[i])
		}
	}
	return results
}



func main(){
	fmt.Println(uncommonFromSentences("apple apple","banana"))
}