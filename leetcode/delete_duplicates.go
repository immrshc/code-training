package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates1(head *ListNode) *ListNode {
	// -100 <= Node.val <= 100
	dist := make([]int, 201)
	for node := head; node != nil; node = node.Next {
		dist[node.Val+100]++
	}
	var first, last *ListNode
	for i, count := range dist {
		if count != 1 {
			continue
		}
		node := &ListNode{Val: i - 100, Next: nil}
		if first == nil {
			first = node
			last = node
		} else {
			last.Next = node
		}
		last = node
	}
	return first
}

func DeleteDuplicates2(head *ListNode) *ListNode {
	histogram := make(map[int]int)
	var first, last *ListNode
	for node := head; node != nil; node = node.Next {
		histogram[node.Val]++
		if histogram[node.Val] == 1 && (node.Next == nil || node.Val != node.Next.Val) {
			ln := &ListNode{
				Val:  node.Val,
				Next: nil,
			}
			if first == nil {
				first = ln
				last = ln
			} else {
				last.Next = ln
			}
			last = ln
		}
	}
	return first
}
