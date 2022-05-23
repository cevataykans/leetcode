func merge(nums1 []int, m int, nums2 []int, n int)  {
    
    tempList := make( []int, m + n)

    i, j, curI := 0, 0, 0
    for i < m && j < n {
        
        if nums1[ i] > nums2[ j] {
            
            tempList[ curI] = nums2[ j]
            j++
            
        } else {
            
            tempList[ curI] = nums1[ i]
            i++
        }
        curI++
    }
    
    for i < m {
        
        tempList[ curI] = nums1[ i]
        i++
        curI++
    }
    
    for j < n {
        
        tempList[ curI] = nums2[ j]
        j++
        curI++
    }
    
    copy( nums1, tempList)
}