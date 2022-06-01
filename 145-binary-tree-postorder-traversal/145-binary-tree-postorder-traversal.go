/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    
    solution := make( []int, 0)
    return postTraversal( root, solution)
}

func postTraversal( cur *TreeNode, sol []int) []int {
    
    if cur == nil {
        
        return sol
    }
    
    sol = postTraversal( cur.Left, sol)
    sol = postTraversal( cur.Right, sol)
    sol = append( sol, cur.Val)
    return sol
}