package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapExist := make(map[string]int, 0)

	for _, v := range s {
		mapExist[string(v)]++
	}

	for _, v := range t {

		key := string(v)
		if mapExist[key] == 0 {
			return false
		}
		mapExist[key]--

	}
	return true

}

func main() {
	fmt.Println(isAnagram("rat", "car"))
}
