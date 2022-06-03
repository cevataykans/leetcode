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
    
    if root.Val < val {
        
        inserIntoBSTHelper( root, root.Right, val)
    } else {
        inserIntoBSTHelper( root, root.Left, val)   
    }
    return root
}

func inserIntoBSTHelper( prev *TreeNode, cur *TreeNode, val int) {
    
    if cur == nil {
        
        if prev.Val >= val {
            
            prev.Left = new(TreeNode)
            prev.Left.Val = val
        } else {
            
            prev.Right = new( TreeNode)
            prev.Right.Val = val
        }
        return
    }
    
    if cur.Val < val {
        
        inserIntoBSTHelper( cur, cur.Right, val)
    } else {
        inserIntoBSTHelper( cur, cur.Left, val)   
    }
}