package main

import (
	"fmt"
)

func main() {
	l1 := &ListNode{}
	l2 := &ListNode{}
	generate(l1, []int{9, 9, 9, 9, 9, 9})
	generate(l2, []int{9, 9, 9})

	out := addTwoNumbers(l1, l2)

	i := 0
	for out != nil {
		fmt.Println(out.Val, out.Next == nil)
		out = out.Next
		i++
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var s1 []int

	for i := 0;;i++ {
		if l1 == nil && l2 == nil {
			break
		}
		if l1 == nil {
			s1 = append(s1, l2.Val)
			l2 = l2.Next
			continue
		}
		if l2 == nil {
			s1 = append(s1, l1.Val)
			l1 = l1.Next
			continue
		}
		s1 = append(s1, l1.Val + l2.Val)
		l1 = l1.Next
		l2 = l2.Next
	}

	out := &ListNode{}
	generate(out, s1)

	return out
}

func generate(l *ListNode, s []int) {
	if s[0] > 9 && len(s) == 1 {
		s[0] -= 10
		s = append(s, 1)
	} else if s[0] > 9 {
		s[0] -= 10
		s[1]++
	}
	l.Val = s[0]
	if len(s) == 1 {
		return
	}
	l.Next = &ListNode{}
	generate(l.Next, s[1:])
}