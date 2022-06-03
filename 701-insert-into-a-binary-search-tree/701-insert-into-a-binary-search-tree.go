/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    
    if root == nil {
        
        root = &TreeNode{ Val: val, Left: nil, Right: nil }
        return root
    }
    
    inserIntoBSTHelper( root, val)
    return root
}

func inserIntoBSTHelper( cur *TreeNode, val int) {
    
    if cur.Val < val {
        
        if cur.Right != nil {
            
            inserIntoBSTHelper( cur.Right, val)
            
        } else {
            
            cur.Right = new(TreeNode)
            cur.Right.Val = val
        }
        
    } else {
        
        if cur.Left != nil {
            
            inserIntoBSTHelper( cur.Left, val)
            
        } else {
            
            cur.Left = new( TreeNode)
            cur.Left.Val = val
        }
    }
}