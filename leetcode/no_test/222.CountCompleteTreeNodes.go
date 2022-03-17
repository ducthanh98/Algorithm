package main

import "math"

type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
 }

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left,right := root,root
	height := 0

	for left != nil && right != nil{
		left = left.Left
		right = right.Right
		height++

	}
	if left == nil && right == nil {
		return int(math.Pow(2.0,float64(height)) - 1)
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
func main()  {
}