package main

import "github.com/ducthanh98/Algorithm/leetcode/utils"

// dfs
func kthSmallestDst(root *utils.TreeNode, k int) int {
	if root == nil {
		return 0
	}
	var count = 0
	var res = 0
	var dfs func(node *utils.TreeNode)
	dfs = func(node *utils.TreeNode) {
		if node == nil || count == k {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		count++
		if count == k {
			res = k
			return
		}
	}
	return res

}

// bfs
func kthSmallestBst(root *utils.TreeNode, k int) int {

}
