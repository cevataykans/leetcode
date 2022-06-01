/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    
    solution := []int{}
    stack := []*TreeNode{}
    
    cur := root
    for cur != nil || len( stack) > 0 {
        
        for cur != nil {
            
            stack = append( stack, cur)
            cur = cur.Left
        }
        
        cur = stack[ len( stack) - 1]
        stack = stack[ :len( stack) - 1]
        solution = append( solution, cur.Val)
        cur = cur.Right
    }
    return solution
}