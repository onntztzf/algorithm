package main

//94. 二叉树的中序遍历
//https://leetcode.cn/problems/same-tree/
//关键：深度优先搜索或者广度优先搜索

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//深度优先搜索(DFS)
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

//广度优先搜索(BFS)
func isSameTree1(p *TreeNode, q *TreeNode) bool {
	//两棵树都是空，所以相同
	if p == nil && q == nil {
		return true
	}
	//初始化队列
	queue := make([]*TreeNode, 0, 2)
	//添加 p 和 q
	queue = append(queue, p, q)
	for len(queue) > 0 {
		tempP := queue[0]
		tempQ := queue[1]
		//临时节点都是空则判断下一对临时节点
		if tempP == nil && tempQ == nil {
			//丢弃前一对已经对比过的临时节点
			queue = queue[2:]
			continue
		}
		//任意一个临时节点为空或者临时节点的值不想等，则两个树不相同
		if (tempP == nil || tempQ == nil) || (tempP.Val != tempQ.Val) {
			return false
		}
		queue = append(queue, tempP.Left, tempQ.Left, tempP.Right, tempQ.Right)
		//丢弃前一对已经对比过的临时节点
		queue = queue[2:]
	}
	//所有节点都已经对比过，且一致，所以两个树相同
	return true
}
