/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
    charge := 0
    left := l1
    right := l2
    
    var sumList *ListNode = nil
    var head *ListNode
    
    for left != nil && right != nil {
        
        if sumList == nil {
            
            sumList = &ListNode{}
            head = sumList
            
        } else {
            
            sumList.Next = &ListNode{}
            sumList = sumList.Next
        }
        
        sum := left.Val + right.Val + charge
        charge = sum / 10
        sum = sum % 10
        sumList.Val = sum
        
        left = left.Next
        right = right.Next
    }
    
    for left != nil {
        
        sumList.Next = &ListNode{}
        sumList = sumList.Next
        
        sum := left.Val + charge
        charge = sum / 10
        sum = sum % 10
        sumList.Val = sum
        
        left = left.Next
    }
    
    for right != nil {
        
        sumList.Next = &ListNode{}
        sumList = sumList.Next
        
        sum := right.Val + charge
        charge = sum / 10
        sum = sum % 10
        sumList.Val = sum
        
        right = right.Next
    }
    
    if charge != 0 {
        
        sumList.Next = &ListNode{}
        sumList = sumList.Next
        sumList.Val = charge
    }
    
    return head
}