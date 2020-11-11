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