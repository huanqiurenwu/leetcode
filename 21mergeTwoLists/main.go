package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 3, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7}}}
	l2 = mergeTwoLists2(l1, l2)
	print(l2)
}

func print(node *ListNode) {
	for node != nil {
		fmt.Println(node)
		node = node.Next
	}
}
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	head1 := l1
	head2 := l2
	var res *ListNode
	for {
		if l1.Next != nil && l2 != nil {
			if l1.Val < l2.Val && l2.Val <= l1.Next.Val {
				tmp := l2.Next
				l2.Next = l1.Next
				l1.Next = l2
				l1 = l1.Next
				l2 = tmp
			} else if l2.Val > l1.Next.Val {
				l1 = l1.Next
			}
		}
		if l2.Next != nil && l1 != nil {
			if l2.Val <= l1.Val && l1.Val <= l2.Next.Val {
				tmp := l1.Next
				l1.Next = l2.Next
				l2.Next = l1
				l1 = tmp
				l2 = l2.Next
			} else if l1.Val > l2.Next.Val {
				l2 = l2.Next
			} else if l1.Val < l2.Val && l1.Next == nil {
				l1.Next = l2
			}
		}
		if l1 == nil || l2 == nil {
			break
		}
	}
	if head1.Val < head2.Val {
		res = head1
	} else {
		res = head2
	}
	return res
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	var prev *ListNode
	if l1 == nil && l2 == nil {
		return prev
	}
	if l1 == nil && l2 != nil {
		return l2
	}
	if l1 != nil && l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		prev = l1
		l1 = l1.Next
	} else {
		prev = l2
		l2 = l2.Next
	}
	head := prev
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1 != nil {
		prev.Next = l1
	}
	if l2 != nil {
		prev.Next = l2
	}
	return head
}
