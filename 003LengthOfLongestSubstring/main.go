package main

import "fmt"

func main() {
	var test = []struct {
		src    string
		expect int
	}{
		{src: "abcabcbb", expect: 3},
		{src: "abcabcetfbb", expect: 6},
		{src: "abcdef", expect: 6},
		{src: "pwwkew", expect: 3},
		{src: "", expect: 0},
		{src: " ", expect: 1},
	}

	for _, x := range test {
		if got := lengthOfLongestSubstring(x.src); got != x.expect {
			fmt.Println(fmt.Sprintf("test failed ,src: %s, want %d got %d", x.src, x.expect, got))
			return
		}
	}
}

func lengthOfLongestSubstring(s string) int {
	return max(s[0], s[1:], 0)
}

func max(tar byte, src string, n int) int {
	if len(src) == 1 {
		return n
	}
	fmt.Println("new round")

	for i := range src {
		if src[i] == tar {
			if n < i+1 {
				n = i + 1
				fmt.Println("1==", n)
			}
			break
		}
		if n < i+1 {
			n = i + 1
			fmt.Println("2==", n)
		}
	}

	return max(src[0], src[1:], n)
}
