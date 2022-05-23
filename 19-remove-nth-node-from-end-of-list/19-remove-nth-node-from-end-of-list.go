/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
        
    nodeList := make( []*ListNode, 0)
    
    curPtr := head
    for curPtr != nil {
        
        nodeList = append( nodeList, curPtr)
        curPtr = curPtr.Next
    }
    
    n = len( nodeList) - n
    if n == 0 {
        
        head = head.Next
        return head
    }
    
    nodeList[ n - 1].Next = nodeList[ n].Next 
    return nodeList[ 0]
}