package main

import "fmt"

func main()  {
	arr := [10]int{2, 3, 5, 7, 11, 13}

	//data := 6
	position := 10

	for i := position; i <  9 ; i++ {
		arr[i] = arr[i+1]
	}
	fmt.Println(arr)

}
