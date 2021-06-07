package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	arr := strings.Split(path,"/")
	stack := make([]string,0)

	for _,v := range arr {
		if v != "" && v != ".." && v!= "." {

			stack = append(stack,v)

		} else if v != ""{
			countDot := len(v) -1
			idx := len(stack)-countDot
			if idx <0 {
				idx = 0
			}
			stack = stack[:idx]
		}
	}

	result := "/"+strings.Join(stack,"/")

	return result
}

func main()  {
	fmt.Println(simplifyPath("/..."))
}

//