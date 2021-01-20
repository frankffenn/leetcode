# 两数相加

一开始的解法：

```
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var s1 []int

	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}

	for i := 0; l2 != nil; i++ {
		if len(s1) > i {
			s1[i] += l2.Val
		} else {
			s1 = append(s1, l2.Val)
		}

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
```

执行用时: 16 ms
内存消耗: 5.4 MB


优化后减少循环次数，提高效率，代码如下，还能再优化一下内存使用率：
```
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
```


执行用时：8 ms, 在所有 Go 提交中击败了92.56%的用户
内存消耗：5.4 MB, 在所有 Go 提交中击败了5.05%的用户