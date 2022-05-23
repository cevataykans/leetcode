func merge(nums1 []int, m int, nums2 []int, n int)  {
    
    i, j, curI := m - 1, n - 1, m + n - 1
    for i >= 0 && j >= 0 {
        
        if nums1[ i] < nums2[ j] {
            
            nums1[ curI] = nums2[ j]
            j--
            
        } else {
            
            nums1[ curI] = nums1[ i]
            i--
        }
        curI--
    }
    
    for j >= 0 {
        
        nums1[ curI] = nums2[ j]
        j--
        curI--
    }
    
}