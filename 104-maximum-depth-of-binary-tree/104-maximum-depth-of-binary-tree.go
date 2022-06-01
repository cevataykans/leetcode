/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    
    if root == nil {
        
        return 0
    }
    
    leftSub := 1 + maxDepth( root.Left)
    rightSub := 1 + maxDepth( root.Right)
    
    if leftSub > rightSub {
        
        return leftSub
    }
    return rightSub
}