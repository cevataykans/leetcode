func checkStraightLine(coordinates [][]int) bool {
    
    x1, x2, y1, y2 := coordinates[ 0][ 0], coordinates[ 1][ 0], coordinates[ 0][ 1], coordinates[ 1][ 1]
    dx, dy := x2 - x1, y2 - y1
    for i := 0; i < len( coordinates); i++ {
        
        if (coordinates[ i][ 0] - x2) * dy != dx * (coordinates[ i][ 1] - y2) {
            
            return false
        }
    }
    return true
}
