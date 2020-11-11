package list

func mergeTwoSortedList(a, b *ListNode) *ListNode {
	var head *ListNode = &ListNode{}
	prev := head

	for a != nil && b != nil {
		if a.Val < b.Val {
			prev.Next = a
			a = a.Next
		} else {
			prev.Next = b
			b = b.Next
		}
		prev = prev.Next
	}

	if a == nil {
		prev.Next = b
	} else {
		prev.Next = a
	}
	return head.Next
}

func MergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	left := head
	right := slow.Next
	slow.Next = nil

	r1 := MergeSort(left)
	r2 := MergeSort(right)
	return mergeTwoSortedList(r1, r2)
}


// head is a pointer, create a new head with no value
// make it as prev node
func BubbleSort(head *ListNode) *ListNode {
	newHead := &ListNode{0, head}	
	var tail *ListNode = nil
	prev := newHead
	cur := prev.Next
	for newHead.Next != tail {
		for cur.Next != tail {
			if cur.Val > cur.Next.Val {
				prev.Next = cur.Next
				cur.Next = cur.Next.Next
				prev.Next.Next = cur
				prev = prev.Next
			} else {
				cur = cur.Next
				prev = prev.Next
			}
		}	
		tail = cur
		prev = newHead
		cur = prev.Next
	}
	return newHead.Next
}