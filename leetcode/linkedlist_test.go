package leetcode

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	head := &ListNode{
		Val: 1,
	}
	p := head
	for i := 2; i < 5; i++ {
		c := &ListNode{
			Val: i,
		}
		p.Next = c
		p = p.Next
	}

	h := swapPairsRec(head)

	for p = h; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}
}
