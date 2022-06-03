/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    
    _, ok := findHeight( root)
    return ok
}

func findHeight( cur *TreeNode) (int, bool) {
    
    if cur == nil {
        
        return 0, true
    }
    
    leftH, leftOk := findHeight( cur.Left)
    rightH, rightOk := findHeight( cur.Right)
    
    if leftH > rightH {
        
        leftH, rightH = rightH, leftH
    }
    
    if rightH - leftH <= 1 {
        
        return maxInt( leftH, rightH) + 1, leftOk && rightOk
    }
    return -1, false
}

func maxInt( a int, b int) int {
    
    if a > b {
        
        return a
    }
    return b
}