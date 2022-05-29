func canMakeArithmeticProgression(arr []int) bool {
    
    sort.Ints( arr)
    
    for i := 1; i < len( arr) - 1; i++ {
        
        if arr[ i] - arr[ i - 1] != arr[ i + 1] - arr[ i] {
            
            return false
        }
    }
    return true
}