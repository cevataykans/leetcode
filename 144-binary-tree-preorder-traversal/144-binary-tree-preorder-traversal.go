/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    
    solution := []int{}
    return traversePre( solution, root)
}

func traversePre( sol []int, cur *TreeNode) []int {
    
    if cur == nil {
        
        return sol
    }
    
    sol = append( sol, cur.Val)
    sol = traversePre( sol, cur.Left)
    sol = traversePre( sol, cur.Right)
    return sol
}