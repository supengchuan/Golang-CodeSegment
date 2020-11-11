package list

import "fmt"

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintList(l *ListNode) {
	if l == nil {
		fmt.Println("list is nil")
	}
	p := l
	for p != nil {
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
	fmt.Println()
}

func InitAList(nums []int) *ListNode {
	n := len(nums)

	if n == 0 {
		return nil
	}
	head := &ListNode{nums[0], nil}
	p := head
	for i := 1; i < n; i++ {
		p.Next = &ListNode{nums[i], nil}
		p = p.Next
	}

	return head
}

// leetcode 237
func DeleteNode(node *ListNode) {
	nextNode := node.Next
	node.Val = nextNode.Val
	node.Next = nextNode.Next
	nextNode.Next = nil
}

// leetcode  21
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var result *ListNode
	if l1.Val <= l2.Val {
		result = l1
		result.Next = mergeTwoLists(l1.Next, l2)
	} else {
		result = l2
		result.Next = mergeTwoLists(l1, l2.Next)
	}
	return result
}

func MergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}

	left := MergeKLists(lists[0 : n/2])
	right := MergeKLists(lists[n/2 : n])
	return mergeTwoLists(left, right)
}
