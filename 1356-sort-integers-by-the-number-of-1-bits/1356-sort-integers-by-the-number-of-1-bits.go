func sortByBits(arr []int) []int {
    
    sort.SliceStable( arr, func( f int, s int) bool {
        
        fC := cb( arr[f]) * 10000 + arr[f]
        sC := cb( arr[s]) * 10000 + arr[s]
        
        return fC < sC
    })
    return arr
}

func cb( i int) int {
        
    count := 0
    for i != 0 {
        
        i = i & ( i - 1)
        count++
    }
    return count
}