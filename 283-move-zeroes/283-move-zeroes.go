func moveZeroes(nums []int)  {
    
    placeI := 0
    for i := 0; i < len( nums); i++ {
        
        if nums[ i] != 0 {
            
            temp := nums[ placeI]
            nums[ placeI] = nums[ i]
            nums[ i] = temp
            placeI++
        }
    }
}