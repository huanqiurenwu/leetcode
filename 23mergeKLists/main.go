package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//input1 := &ListNode{
	//	Val:1,
	//	Next:&ListNode{
	//		Val:2,
	//		Next:&ListNode{
	//			Val:3,
	//			Next:&ListNode{
	//				Val:4,
	//				Next:&ListNode{
	//					Val:5,
	//				},
	//			},
	//		},
	//	},
	//}
	//input2 := &ListNode{
	//	Val:1,
	//	Next:&ListNode{
	//		Val:2,
	//		Next:&ListNode{
	//			Val:3,
	//			Next:&ListNode{
	//				Val:4,
	//				Next:&ListNode{
	//					Val:5,
	//				},
	//			},
	//		},
	//	},
	//}
	//input3 := &ListNode{
	//	Val:1,
	//	Next:&ListNode{
	//		Val:2,
	//		Next:&ListNode{
	//			Val:3,
	//			Next:&ListNode{
	//				Val:4,
	//				Next:&ListNode{
	//					Val:5,
	//				},
	//			},
	//		},
	//	},
	//}
	inputs := []*ListNode{&ListNode{}, &ListNode{}, &ListNode{}}
	output := mergeKLists(inputs)
	for head := output; head != nil; head = head.Next {
		fmt.Print("  ", head.Val)
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 1 {
		return lists[0]
	} else if len(lists) == 0 {
		return nil
	}
	mid := len(lists) / 2
	resultL := mergeKLists(lists[:mid])
	resultR := mergeKLists(lists[mid:])
	return merge2List(resultR, resultL)
}

func merge2List(listA, listB *ListNode) *ListNode {
	if listA == nil {
		return listB
	} else if listB == nil {
		return listA
	}
	var head *ListNode
	if listA.Val < listB.Val {
		head = listA
		listA = listA.Next
	} else {
		head = listB
		listB = listB.Next
	}
	var tmp *ListNode
	for tmp = head; listA != nil && listB != nil; tmp = tmp.Next {
		if listA.Val < listB.Val {
			tmp.Next = listA
			listA = listA.Next
		} else {
			tmp.Next = listB
			listB = listB.Next
		}
	}
	if listA == nil {
		tmp.Next = listB
	} else {
		tmp.Next = listA
	}
	return head
}
