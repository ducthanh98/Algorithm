package main

import "fmt"

func getIthBit(n, i int) int {
	mask := 1 << i
	result := n & mask
	if result > 0 {
		return 1
	} else {
		return 0
	}
}

func clearIthBit(n, i int) int {
	mask := ^(1 << i)
	result := n & mask
	return result

}

func clearLastBits(n, i int) int {
	mask := -1 << i
	result := n & mask
	return result

}

func setIthBit(n, i int) int {
	mask := 1 << i
	n = n | mask
	return n
}

func updateIthBit(n, i, v int) int {
	n = clearIthBit(n, i)
	mask := v << i
	n = n | mask
	return n
}

func main() {
	fmt.Println(clearLastBits(15, 2))
}
