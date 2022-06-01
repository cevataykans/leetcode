type P struct {
    
    I int
    J int
}

func updateMatrix(mat [][]int) [][]int {
    
    grid := make( [][]int, len( mat), len(mat))
    for i := 0; i < len( grid); i++ {
        
        grid[ i] = make( []int, len( mat[ i]), len( mat[ i]) )
    }
    
    queue := make( []P, 0)
    for i := 0; i < len( mat); i++ {
        for j := 0; j < len( mat[ i]); j++ {
            
            if mat[ i][ j] == 0 {
                
                queue = append( queue, P{ I: i, J: j})
            }
        }
    }
    
    distance := 0
    visited := make( map[ P]bool)
    for len( queue) > 0 {
        
        toVisit := queue
        queue = make( []P, 0)
        for len( toVisit) > 0 {
        
            curP := toVisit[ 0]
            toVisit = toVisit[ 1:]
            if curP.I < 0 || curP.I >= len( mat) || curP.J < 0 || curP.J >= len( mat[ 0]) || visited[ curP] {
                
                continue
            }
            
            visited[ curP] = true
            if mat[ curP.I][ curP.J] == 1 {
                
                grid[ curP.I][ curP.J] = distance
            }
            
            queue = append( queue, P{ I: curP.I + 1, J: curP.J })
            queue = append( queue, P{ I: curP.I - 1, J: curP.J })
            queue = append( queue, P{ I: curP.I, J: curP.J - 1 })
            queue = append( queue, P{ I: curP.I, J: curP.J + 1 })
        }
        distance++
    }
    
    return grid
}