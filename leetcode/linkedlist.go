package leetcode

// ListNode - Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/submissions/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := ListNode{
		Next: head,
	}
	// use dummy to unify the
	p1 := &dummy
	p2 := &dummy
	for i := 0; i < n+1; i++ {
		p1 = p1.Next
	}

	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	p2.Next = p2.Next.Next
	return dummy.Next
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	p := dummy

	for p != nil {
		p1 := p
		p = p.Next
		if p == nil {
			break
		}
		p2 := p.Next

		tmp := p2.Next
		p1.Next = p2
		p2.Next = p
		p.Next = tmp
	}

	return dummy.Next
}

func swapPairsRec(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// to swap head and head.Next, next will be the new head
	next := head.Next
	head.Next = swapPairsRec(next.Next)
	next.Next = head
	return next
}
