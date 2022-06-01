func sumOddLengthSubarrays(arr []int) int {
    
    sum := 0
    for i := 0; i < len( arr); i++ {
        
        sum += arr[ i] * int( math.Ceil( float64(i + 1) * float64( len( arr) - i) / 2 ) )
    }
    return sum
}