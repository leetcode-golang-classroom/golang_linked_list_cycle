package sol

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	slow := head
	if slow == nil {
		return false
	}
	fast := head.Next
	if fast == nil {
		return false
	}
	for slow != nil && fast != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}

	return false
}
