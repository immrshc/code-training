package leetcode

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	first := &ListNode{
		Val:  0,
		Next: nil,
	}
	last := first
	var carry int
	// 最後はcarryだけを追加して終了する
	for l1 != nil || l2 != nil || carry != 0 {
		var x, y int
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		carry = sum / 10
		if sum >= 10 {
			sum = sum % 10
		}
		last.Next = &ListNode{
			Val:  sum,
			Next: nil,
		}
		last = last.Next
	}
	return first.Next
}
