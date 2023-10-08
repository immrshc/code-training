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
	// -100 <= Node.val <= 100
	pseudo := &ListNode{Val: -101, Next: nil}
	first, last := pseudo, pseudo
	for i, count := range dist {
		if count != 1 {
			continue
		}
		node := &ListNode{Val: i - 100, Next: nil}
		last.Next = node
		last = last.Next
	}
	return first.Next
}

func DeleteDuplicates2(head *ListNode) *ListNode {
	histogram := make(map[int]int)
	// -100 <= Node.val <= 100
	pseudo := &ListNode{Val: -101, Next: nil}
	first, last := pseudo, pseudo
	for node := head; node != nil; node = node.Next {
		histogram[node.Val]++
		if histogram[node.Val] == 1 && (node.Next == nil || node.Val != node.Next.Val) {
			ln := &ListNode{
				Val:  node.Val,
				Next: nil,
			}
			last.Next = ln
			last = last.Next
		}
	}
	return first.Next
}

func DeleteDuplicates3(head *ListNode) *ListNode {
	// -100 <= Node.val <= 100
	pseudo := &ListNode{Val: -101, Next: nil}
	first, last, pred := pseudo, pseudo, pseudo
	for node := head; node != nil; node = node.Next {
		// ソートされており、重複がなければ前後と値が異なる
		// ... → 1 → 2 → 3 → ...
		if node.Next == nil || node.Val != node.Next.Val {
			if pred.Val != node.Val {
				next := &ListNode{
					Val:  node.Val,
					Next: nil,
				}
				last.Next = next
				last = last.Next
			}
		}
		pred = node
	}
	return first.Next
}
