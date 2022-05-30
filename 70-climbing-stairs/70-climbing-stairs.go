func climbStairs(n int) int {
    
    steps := make( []int, n + 2, n + 2)
    steps[ 1] = 1
    steps[ 2] = 2
    for i := 3; i <= n; i++ {
        
        steps[ i] = steps[ i - 1] + steps[ i - 2]
    }
    return steps[ n]
}