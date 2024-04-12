package main

import (
	"fmt"
	"strings"
)

func findDuplicate(paths []string) [][]string {
	mapContentDirectory := make(map[string][]string, 0)
	for _, path := range paths {
		arr := strings.Split(path, " ")
		for i := 1; i < len(arr); i++ {
			idx := strings.Index(arr[i], "(")
			directory, content := arr[i][:idx], arr[i][idx+1:len(arr[i])-1]
			if mapContentDirectory[content] == nil {
				mapContentDirectory[content] = []string{fmt.Sprintf("%v/%v", arr[0], directory)}
			} else {
				mapContentDirectory[content] = append(mapContentDirectory[content], fmt.Sprintf("%v/%v", arr[0], directory))
			}
		}
	}
	results := make([][]string, 0)
	for _, val := range mapContentDirectory {
		if len(val) > 1 {
			results = append(results, val)
		}
	}
	return results
}
