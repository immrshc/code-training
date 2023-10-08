package leetcode

/*
 Definition for singly-linked list.
 type ListNode struct {
     Val int
     Next *ListNode
 }
*/

func MergeTwoLists1(list1, list2 *ListNode) *ListNode {
	pseudo := &ListNode{
		Val:  0,
		Next: nil,
	}
	last := pseudo
	for list1 != nil && list2 != nil {
		var next *ListNode
		if list1.Val < list2.Val {
			next = &ListNode{
				Val:  list1.Val,
				Next: nil,
			}
			list1 = list1.Next
		} else {
			next = &ListNode{
				Val:  list2.Val,
				Next: nil,
			}
			list2 = list2.Next
		}
		last.Next = next
		last = next
	}
	if list1 == nil {
		last.Next = list2
	}
	if list2 == nil {
		last.Next = list1
	}
	return pseudo.Next
}

func MergeTwoLists2(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = MergeTwoLists2(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoLists2(list1, list2.Next)
		return list2
	}
}
