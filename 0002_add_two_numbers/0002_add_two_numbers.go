package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode
	var l3Current *ListNode
	tl1 := l1
	tl2 := l2
	carry := 0

	for {
		if tl1 == nil && tl2 == nil && carry == 0 {
			break
		}

		val := carry
		if tl1 != nil {
			val += tl1.Val
		}
		if tl2 != nil {
			val += tl2.Val
		}

		diff := val - 10
		if diff >= 0 {
			val = diff
			carry = 1
		} else {
			carry = 0
		}

		node := &ListNode{
			Val:  val,
			Next: nil,
		}

		if tl1 != nil {
			tl1 = tl1.Next
		}
		if tl2 != nil {
			tl2 = tl2.Next
		}

		if l3 == nil {
			l3 = node
			l3Current = l3
		} else {
			l3Current.Next = node
			l3Current = node
		}
	}

	return l3
}
