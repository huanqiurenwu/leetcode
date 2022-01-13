package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head = &ListNode{
		Next: head,
	}
	for prev, cur, next := head, head.Next, head.Next.Next.Next; ; {
		prev.Next = cur.Next
		cur.Next.Next = cur
		cur.Next = next
		if next == nil || next.Next == nil {
			break
		}
		prev = cur
		cur = next
		next = next.Next.Next
	}
	return head.Next
}

func main() {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	output := swapPairs(input)
	for head := output; head != nil; head = head.Next {
		fmt.Print("  ", head.Val)
	}
}
