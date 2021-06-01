/*
	83. 删除排序链表中的重复元素
	https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
 */
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	currentNode := head
	for currentNode.Next != nil {
		nextNode := currentNode.Next
		if currentNode.Val == nextNode.Val {
			currentNode.Next = nextNode.Next
		} else {
			currentNode = nextNode
		}
	}
	return head
}

func main() {

}
