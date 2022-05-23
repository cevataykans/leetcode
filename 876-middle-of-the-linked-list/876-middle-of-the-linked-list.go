/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    
    nodeList := make( []*ListNode, 0)
    
    curPtr := head
    for curPtr != nil {
        
        nodeList = append( nodeList, curPtr)
        curPtr = curPtr.Next
    }
    
    return nodeList[ len( nodeList) / 2]
}