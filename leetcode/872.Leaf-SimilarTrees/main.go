package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	r1, r2 := make([]int, 0), make([]int, 0)
	dfs(root1, &r1)
	dfs(root2, &r2)
	fmt.Println(r1, r2)
	return reflect.DeepEqual(r1, r2)
}

func dfs(root *TreeNode, r *[]int) {
	if root.Left == nil && root.Right == nil {
		*r = append(*r, root.Val)
		return
	}
	dfs(root.Left, r)
	dfs(root.Right, r)
}

func test(a *[]int) {
	*a = append(*a, 1)
}

func main() {
	a := make([]int, 0)
	test(&a)
	fmt.Println(a)
}
