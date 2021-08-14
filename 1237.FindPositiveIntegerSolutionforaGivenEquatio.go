package main

import "fmt"

/**
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */

func findSolution(customFunction func(int, int) int, z int) [][]int {
	results := make([][]int, 0)
	for i := 1; i <= 1000; i++ {
		for j := 1; j <= 1000; j++ {

			tmp := customFunction(i, j)
			if tmp < z {
				break
			} else if z == tmp {
				results = append(results, []int{i, j})
			}

		}

	}
	return results
}

func main() {

	fmt.Println(findSolution(func(i1, i2 int) int { return i1 + i2 }, 5))
}
