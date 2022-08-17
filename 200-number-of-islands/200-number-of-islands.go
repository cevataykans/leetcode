type P struct {
    I int
    J int
}

func numIslands(grid [][]byte) int {
    
    numOfIslands := 0
    visited := make(map[P]bool)
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[ i]); j++ {
            curP := P{ I: i, J: j}
            if !visited[ curP ] && grid[ i][ j] == '1' {
                numOfIslands++
                visit( grid, curP, visited)
            }
        }
    }
    return numOfIslands
}

func visit( grid [][]byte, p P, visited map[P]bool) {
    
    if p.I < 0 || p.I == len(grid) || p.J < 0 || p.J == len(grid[0]) {
        return
    }
    
    if visited[ p] || grid[ p.I][p.J] == '0' {
        return
    } 
    
    visited[ p] = true
    visit( grid, P{ I: p.I + 1, J: p.J}, visited)
    visit( grid, P{ I: p.I - 1, J: p.J}, visited)
    visit( grid, P{ I: p.I, J: p.J - 1}, visited)
    visit( grid, P{ I: p.I, J: p.J + 1}, visited)
}