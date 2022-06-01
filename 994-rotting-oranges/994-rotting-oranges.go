type P struct {
    
    I int
    J int
}

func orangesRotting(grid [][]int) int {
    
    visited := make( map[ P]bool)
    queue := []P{}
    freshCount := 0
    for i := 0; i < len( grid); i++ {
        
        for j := 0; j < len( grid[ i]); j++ {
            
            if grid[ i][ j] == 2 {
                
                queue = append( queue, P{ I: i, J: j })
                
            } else if grid[ i][ j] == 1 {
                
                freshCount++
            }
        }
    }
    
    minute := 0
    for len( queue) > 0 {
                
        toVisit := queue
        queue = []P{}
        
        for _, point := range toVisit {
            
            if !inRange( point, grid) || visited[ point] || grid[ point.I][ point.J] == 0 {
                
                continue
            }
            
            visited[ point] = true
            
            if grid[ point.I][ point.J] == 1 {
                
                freshCount--
            }
            
            queue = append( queue, P{ I: point.I + 1, J: point.J})
            queue = append( queue, P{ I: point.I - 1, J: point.J})
            queue = append( queue, P{ I: point.I, J: point.J + 1})
            queue = append( queue, P{ I: point.I, J: point.J - 1})
        }
        
        if freshCount == 0 {
            
            return minute
        }
        minute++
    }
    
    if freshCount == 0 {
            
        return minute
    }
    return -1
}

func inRange( p P, grid [][]int ) bool {
    
    return p.I >= 0 && p.J >= 0 && p.I < len( grid) && p.J < len( grid[ 0])
}