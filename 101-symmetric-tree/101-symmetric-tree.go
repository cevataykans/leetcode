/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
 
    return isSymm( root.Left, root.Right)
}

func isSymm( left *TreeNode, right *TreeNode) bool {
    
    if left == nil || right == nil {
        
        return left == right
    }

    if left.Val != right.Val {
        
        return false
    }
    
    return isSymm( left.Left, right.Right) && isSymm( left.Right, right.Left)
}

