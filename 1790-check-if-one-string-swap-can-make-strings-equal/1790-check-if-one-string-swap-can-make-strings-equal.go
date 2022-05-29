func areAlmostEqual(s1 string, s2 string) bool {
    
    if len( s1) != len( s2) {
        
        return false
    }
    
    indices := make( []int, 0)
    for i := 0; i < len(s1); i++ {
        
        if s1[ i] != s2[ i] {
            
            indices = append( indices, i)
        }
        
        if len( indices) > 2 {
            return false
        }
    }
    
    if len( indices) == 0 {
        return true
    }
    
    if len( indices) == 2 {
        
        s1Bytes := make( []byte, len(s1))
        copy( s1Bytes, s1)
    
        s1Bytes[ indices[ 0]], s1Bytes[ indices[ 1]] = s1Bytes[ indices[ 1]], s1Bytes[ indices[ 0]]
        return string( s1Bytes) == s2   
    }
    return false
}