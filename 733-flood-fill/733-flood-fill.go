type Point struct {
    
    C, R int
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    
    queue := []Point{ Point{ C:sc, R:sr } }
    visited := make( map[ Point]bool)
    
    flagColor := image[ sr][ sc]
    for len( queue) > 0 {
        
        curPoint := queue[ 0]
        queue = queue[ 1:]
        
        if visited[ curPoint] {
            
            continue
        }
        visited[ curPoint] = true
        
        if image[ curPoint.R][ curPoint.C] == flagColor {
            
            image[ curPoint.R][ curPoint.C] = newColor
            
            if curPoint.R + 1 < len( image) {

                queue = append( queue, Point{ C: curPoint.C, R: curPoint.R + 1})
            }

            if curPoint.R - 1 >= 0 {

                queue = append( queue, Point{ C: curPoint.C, R: curPoint.R - 1})
            }

            if curPoint.C + 1 < len( image[ 0]) {

                queue = append( queue, Point{ C: curPoint.C + 1, R: curPoint.R})
            }

            if curPoint.C - 1 >= 0 {

                queue = append( queue, Point{ C: curPoint.C - 1, R: curPoint.R})
            }
        }
    }
    return image
}