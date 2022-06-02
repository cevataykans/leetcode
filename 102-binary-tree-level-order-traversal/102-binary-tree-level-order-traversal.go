/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    
    queue := []*TreeNode{}
    solution := [][]int{}
    
    if root != nil {
        
        queue = append( queue, root)
    }
    
    for len( queue) > 0 {
        
        toVisit := queue
        queue = make( []*TreeNode, 0)
        solution = append( solution, make( []int, 0))
        
        for len( toVisit) > 0 {
            
            cur := toVisit[ 0]
            toVisit = toVisit[ 1:]
            
            solution[ len( solution) - 1] = append( solution[ len( solution) - 1], cur.Val)
            
            if cur.Left != nil {
             
                queue = append( queue, cur.Left)
            }
            
            if cur.Right != nil {
                
                queue = append( queue, cur.Right)   
            }
        }
    }
    return solution
}