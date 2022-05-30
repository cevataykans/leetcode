/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    
    set := make( map[*ListNode]bool)
    
    for head != nil {
        
        if set[ head] {
            
            return true
        }
        set[ head] = true
        head = head.Next
    }
    return false
}