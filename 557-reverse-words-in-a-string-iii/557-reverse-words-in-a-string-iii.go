func reverseWords(s string) string {
    
    strList := strings.Split( s, " ")
    for i, word := range strList {
        
        strList[ i] = reverseStr( []byte( word) )
    }
    return strings.Join( strList, " ")
}

func reverseStr( s []byte) string {
    
    i, j := 0, len( s) - 1
    for i < j {
        
        s[ i], s[ j] = s[ j], s[ i]
        i++
        j--
    }
    return string( s)
}