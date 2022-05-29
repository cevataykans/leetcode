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
    
    var mergedTree *TreeNode = new(TreeNode)
    merge( root1, root2, mergedTree)
    return mergedTree
}

func merge( root1 *TreeNode, root2 *TreeNode, newRoot *TreeNode) {
    
    if root1 == nil && root2 == nil {
        
        return
    }
    
    var root1Val int
    if root1 != nil {
        root1Val = root1.Val
    }
    
    var root2Val int
    if root2 != nil {
        root2Val = root2.Val
    }
    
    newRoot.Val = root1Val + root2Val
    
    left1 := getLeft( root1)
    left2 := getLeft( root2)
    if left1 != nil || left2 != nil {
        
        newRoot.Left = new( TreeNode)
        merge( left1, left2, newRoot.Left)
    }
    
    right1 := getRight( root1)
    right2 := getRight( root2)
    if right1 != nil || right2 != nil {
        
        newRoot.Right = new( TreeNode)
        merge( right1, right2, newRoot.Right)
    }
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