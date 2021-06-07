package main

import "fmt"

func reverseString(str string,idx,length int)  string{

	if idx < length {
		char := string(str[idx])
		idx++
		return reverseString(str,idx,length) + char
	} else {
		return ""
	}
}

func main()  {
	str := "what your name"
	length := len(str)
	fmt.Println(reverseString(str,0,length))
}
