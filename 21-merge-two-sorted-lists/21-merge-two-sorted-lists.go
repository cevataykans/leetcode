/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    
    if list1 == nil && list2 == nil {
        
        return nil
    }
    
    var merged *ListNode = nil
    var toReturn *ListNode = nil
    for list1 != nil && list2 != nil {
        
        if merged == nil {
            
            merged = new( ListNode)
            toReturn = merged
        } else {
            
            merged.Next = new( ListNode)
            merged = merged.Next   
        }
        
        if list1.Val > list2.Val {
            
            merged.Val = list2.Val
            list2 = list2.Next
        } else {
            
            merged.Val = list1.Val
            list1 = list1.Next
        }
    }
    
    for list1 != nil {
        
        if merged == nil {
            
            merged = new( ListNode)
            toReturn = merged
        } else {
            
            merged.Next = new( ListNode)
            merged = merged.Next   
        }
        
        merged.Val = list1.Val
        list1 = list1.Next
    }
    
    for list2 != nil {
        
        if merged == nil {
            
            merged = new( ListNode)
            toReturn = merged
        } else {
            
            merged.Next = new( ListNode)
            merged = merged.Next   
        }
        
        merged.Val = list2.Val
        list2 = list2.Next
    }
    
    return toReturn
}