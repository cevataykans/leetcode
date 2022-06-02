type P struct{
    
    I int
    J int
}

func isValidSudoku(board [][]byte) bool {
    
    points := make( map[byte][]P )
    
    for i := 0; i < len( board); i++ {
        
        for j := 0; j < len( board[ i]); j++ {
            
            if board[ i][ j] == '.' {
                
                continue
            }
            
            if points[ board[ i][ j]] == nil {
                
                points[ board[ i][ j]] = []P{}
            }
            
            points[ board[ i][ j]] = append( points[ board[ i][ j]], P{ I: i, J: j } )
        }
    }
    
    for _, list := range points {
        
        for i := 0; i < len( list) - 1; i++ {
            
            for j := i + 1; j < len( list); j++ {
                
                if !IsValid( &list[ i], &list[ j]) {
                    
                    return false
                }
            }
        }
    }
    return true
}

func IsValid( p1 *P, p2 *P) bool {
    
    return p1.I != p2.I && p1.J != p2.J && !(p1.I / 3 == p2.I / 3 && p1.J / 3 == p2.J / 3)
}