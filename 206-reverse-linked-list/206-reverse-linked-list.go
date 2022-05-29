/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    
    if head == nil {
        
        return nil
    }    
    
    return reverse( head)
}

func reverse( cur *ListNode) *ListNode {
    
    if cur.Next == nil {
        
        return cur
    }
    
    newHead := reverse( cur.Next)
    cur.Next.Next = cur
    cur.Next = nil
    return newHead
}