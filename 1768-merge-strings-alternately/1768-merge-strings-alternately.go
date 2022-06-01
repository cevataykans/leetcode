func mergeAlternately(word1 string, word2 string) string {
    
    merged := make( []byte, len( word1) + len( word2))
    i, j, k := 0, 0, 0
    for i < len( word1) && j < len( word2) {
        
        merged[ k] = word1[ i]
        k++
        merged[ k] = word2[ j]
        k++
        i++
        j++
    }
    
    if i < len( word1) {
        
        copy( merged[ k:], word1[ i:])
    }
    
    if j < len( word2) {
        
        copy( merged[ k:], word2[ j:])
    }
    return string( merged)
}