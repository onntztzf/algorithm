package main

/**
 * Definition for a binary tree node.
 *
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//94. 二叉树的中序遍历
//https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
//关键：按照访问左子树——根节点——右子树的方式遍历
//递归
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0, 100)
	var f func(node *TreeNode)
	f = func(node *TreeNode) {
		//如果当前节点未空，则返回
		if node == nil {
			return
		}
		f(node.Left)
		result = append(result, node.Val)
		f(node.Right)
	}
	f(root)
	return result
}

//迭代
func inorderTraversal1(root *TreeNode) []int {
	result := make([]int, 0, 100)
	//自己维护整遍历过程
	stack := make([]*TreeNode, 0, 100)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)
		root = root.Right
	}
	return result
}
