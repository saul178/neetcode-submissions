/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// func reverseList(head *ListNode) *ListNode {
//    	var prev *ListNode = nil
// 	curr := head
// 	for curr != nil {
// 		tmpNxt := curr.Next
// 		curr.Next = prev
// 		prev = curr
// 		curr = tmpNxt
// 	}
// 	return prev
// }
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := head
	if head.Next != nil {
		newHead = reverseList(head.Next)
		head.Next.Next = head
	}
	head.Next = nil
	return newHead
}