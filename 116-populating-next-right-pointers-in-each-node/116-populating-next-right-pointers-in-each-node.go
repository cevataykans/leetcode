/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	
    traverse( root, nil)
    return root
}

func traverse( cur *Node, parent *Node) {
    
    if cur == nil {
        
        return
    }
    
    if cur.Left != nil {
        
        cur.Left.Next = cur.Right
    }
    
    if parent != nil && cur == parent.Right && parent.Next != nil {
        
        cur.Next = parent.Next.Left
    }
    traverse( cur.Left, cur)
    traverse( cur.Right, cur)
}