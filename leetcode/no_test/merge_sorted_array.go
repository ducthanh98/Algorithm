package main

import "fmt"

func mergeSortedArray(a1 []int, a2 []int) []int {
	result := make([]int, 0)
	lengthArr1 := len(a1)
	lengthArr2 := len(a2)

	i, j := 0, 0

	for i < lengthArr1 || j < lengthArr2 {
		if j >= lengthArr2 {
			result = append(result, a1[i])
			i++
		} else if i >= lengthArr1 {
			result = append(result, a2[j])
			j++
		} else if a1[i] < a2[j] {
			result = append(result, a1[i])
			i++
		} else {
			result = append(result, a2[j])
			j++
		}

	}

	return result
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	a1 := nums1[0:m]
	a2 := nums2

	result := make([]int, 0)
	lengthArr1 := len(a1)
	lengthArr2 := len(a2)

	i, j := 0, 0

	for i < lengthArr1 || j < lengthArr2 {
		if j >= lengthArr2 {
			result = append(result, a1[i])
			i++
		} else if i >= lengthArr1 {
			result = append(result, a2[j])
			j++
		} else if a1[i] < a2[j] {
			result = append(result, a1[i])
			i++
		} else {
			result = append(result, a2[j])
			j++
		}

	}

}

func main() {
	fmt.Println(mergeSortedArray([]int{1, 3, 5}, []int{2, 2, 4, 6}))
}
