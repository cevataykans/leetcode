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

func traverse( cur *Node, next *Node) {
    
    if cur == nil {
        
        return
    }
    
    cur.Next = next
    

    traverse( cur.Left, cur.Right)
    if cur.Next != nil {
        traverse( cur.Right, cur.Next.Left)
    } else {
        traverse( cur.Right, nil)   
    }
}