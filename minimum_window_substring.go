package main

import (
	"fmt"
	"strings"
)

func minWindow(s string, t string) string {
	mapValue := make(map[string]int,0)
	chars := strings.Split(t,"")
	parents := strings.Split(s,"")
	for _, char := range chars {
		mapValue[char]++
	}

	start,end,minStart,minLength,counter := 0,0,0,99999999,len(t)

	for end < len(s) {
		if mapValue[parents[end]] > 0 {

			counter--

		}
		mapValue[parents[end]]--
		end++

		for counter == 0 {
			if end - start < minLength {
				minStart = start
				minLength = end - start
			}
			fmt.Println(mapValue[parents[start]])
			mapValue[parents[start]]++
			if mapValue[parents[start]] > 0 {
				counter++
			}

			start++
		}

	}
	if minLength == 99999999 {
		return ""
	} else{
		return strings.Join(parents[minStart:minStart+minLength],"")
	}
}

func main()  {
	fmt.Println(minWindow("ADOBECODEBANC","ABC"))
}