/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList(head *ListNode) *ListNode {
    dump := &ListNode {Val: 0, Next: nil}
    cur := head
    for cur != nil {
        tmp := cur
        cur = cur.Next

        dhead := dump.Next
        dump.Next = tmp
        tmp.Next = dhead
    }

    return dump.Next
}

 func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	reversedListHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return reversedListHead
}

// Iterative version
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head

    for curr != nil {
        tmp := curr.Next
        curr.Next = prev
        prev = curr
        curr = tmp
    }

    return prev
}