/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
    
    var sol int = (1 << 31) - 1
    prev := &TreeNode{ Val: -1 }
    findMinDiffInBST( root, prev, &sol)
    return sol
}

func findMinDiffInBST( cur, prev *TreeNode, sol *int) {
    
    if cur == nil {
        return
    }

    findMinDiffInBST( cur.Left, prev, sol)

    if prev.Val != -1 {
        
        diff := cur.Val - prev.Val
        if diff < *sol {
            *sol = diff
        }
    }
    prev.Val = cur.Val
    
    findMinDiffInBST( cur.Right, prev, sol)
}