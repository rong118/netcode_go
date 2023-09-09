/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dump := &ListNode {Val: 0, Next: nil}
    dump.Next = head

    fast := dump
    for i := 0; i < n; i++ {
        fast = fast.Next;
    }
    
    slow := dump
    for fast.Next != nil {
        slow = slow.Next
        fast = fast.Next
    }

    slow.Next = slow.Next.Next

    return dump.Next
}