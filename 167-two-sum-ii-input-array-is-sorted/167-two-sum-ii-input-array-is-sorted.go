func twoSum(numbers []int, target int) []int {
   
    leftI, rightI := 0, len( numbers) - 1
    for numbers[ leftI] + numbers[ rightI] != target {
        
        if numbers[ leftI] + numbers[ rightI] > target {
            
            rightI--
            
        } else {
            
            leftI++
        }
    }
    return []int{ leftI + 1, rightI + 1 }
}