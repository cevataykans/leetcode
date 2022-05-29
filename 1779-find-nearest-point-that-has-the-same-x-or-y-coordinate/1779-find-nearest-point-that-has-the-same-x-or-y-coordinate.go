func nearestValidPoint(x int, y int, points [][]int) int {
    
    solIndex := -1
    curDistance := math.MaxInt32
    
    for i, curPoint := range points {
        
        if curPoint[ 0] == x || curPoint[ 1] == y {
            
            dist := manhattan( x, curPoint[ 0]) + manhattan( y, curPoint[ 1])
            if dist < curDistance {
                
                curDistance = dist
                solIndex = i
            }
        }
    }
    return solIndex
}

func manhattan( a, b int) int {
    
    num := a - b
    if num < 0 {
        
        return -num
    }
    return num
}