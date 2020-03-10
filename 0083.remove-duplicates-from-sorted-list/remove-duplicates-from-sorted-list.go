package problem0083

// ListNode struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	cursor := head

	for cursor != nil {
		if cursor.Next != nil && cursor.Val == cursor.Next.Val {
			cursor.Next = cursor.Next.Next

			continue
		}

		cursor = cursor.Next
	}

	return head
}
