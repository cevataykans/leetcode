func maxProduct(words []string) int {
    
    allBits := make( []uint, len(words) )
    for i, word := range words {
        
        allBits[ i] = 0
        for _, char := range word {
            
            allBits[ i] |= (1 << (char - 'a') )
        }
    }
    
    var maxProduct int = 0
    for i := 0; i < len( words) - 1; i++ {
        for j := i + 1; j < len( words); j++ {
         
            if allBits[ i] & allBits[ j] == 0 && len( words[ i]) * len( words[ j]) > maxProduct {
                
                maxProduct = len( words[ i]) * len( words[ j])
            } 
        }
    }
    return maxProduct
}