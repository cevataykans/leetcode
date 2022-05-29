/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    
    if root1 == nil && root2 == nil {
        
        return nil
    }

    var curNode TreeNode
    
    var root1Val int
    if root1 != nil {
        root1Val = root1.Val
    }
    
    var root2Val int
    if root2 != nil {
        root2Val = root2.Val
    }
    
    curNode.Val = root1Val + root2Val
    curNode.Left = mergeTrees( getLeft( root1), getLeft( root2))
    curNode.Right = mergeTrees( getRight( root1), getRight( root2))
    return &curNode
}

func getLeft( cur *TreeNode) *TreeNode {
    
    if cur == nil || cur.Left == nil {
        
        return nil
    }
    return cur.Left
}

func getRight( cur * TreeNode) *TreeNode {
    
    if cur == nil || cur.Right == nil {
        
        return nil
    }
    return cur.Right
}