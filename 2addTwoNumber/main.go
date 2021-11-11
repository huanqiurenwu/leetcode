package main

import "fmt"

func main() {
	l1 := ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}
	l2 := ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0}}}
	result := addTwoNumbers(&l1, &l2)
	for ; result != nil; result = result.Next {
		fmt.Printf(" %d", result.Val)
	}
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	point1 := l1
	point2 := l2
	jinwei := 0
	for ; ; point1, point2 = point1.Next, point2.Next {
		tmpVal := point1.Val + point2.Val + jinwei
		jinwei = 0
		if tmpVal < 10 {
			point1.Val = tmpVal
		} else {
			point1.Val = tmpVal - 10
			jinwei = 1
		}
		if point1.Next == nil && point2.Next == nil && jinwei == 0 {
			break
		}
		if point1.Next == nil {
			point1.Next = &ListNode{
				Val: 0,
			}
		}
		if point2.Next == nil {
			point2.Next = &ListNode{
				Val: 0,
			}
		}
	}
	return l1
}
