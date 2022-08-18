func merge(intervals [][]int) [][]int {
   
    merged := make( [][]int, 0)
    sortIntervals( intervals)
    
    merged = append( merged, intervals[ 0])
    cur := 0
    for i := 1; i < len(intervals); i++ {
        if merged[ cur][ 1] >= intervals[ i][ 0]{
            merged[ cur][ 1] = max(merged[ cur][ 1], intervals[ i][ 1])
        } else {
            merged = append( merged, intervals[ i])
            cur++
        }
    }
    
    return merged
}

func sortIntervals( intervals [][]int) {
        sort.SliceStable( intervals, func(i, j int) bool {
        if intervals[ i][ 0] < intervals[ j][ 0] {
            return true
        }
        
        if intervals[ i][ 0] > intervals[ j][ 0] {
            return false
        }
        
        return intervals[ i][ 1] < intervals[ j][ 1]
    })
}

func max( i, j int) int {
    if i > j {
        return i
    }
    return j
}