func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    
    m := len(nums1)
    n := len(nums2)
    l := (m + n + 1) / 2
    r := (m + n + 2) / 2
    num1 := getKth( nums1, nums2, 0, 0, l)
    if l != r {
        num2 := getKth( nums1, nums2, 0, 0, r)
        return float64( num1 + num2) / 2.0
    }
    return float64(num1)
}

func getKth( a, b []int, aStart, bStart, k int ) int {
    
    if aStart >= len(a) {
        return b[ bStart + k - 1]
    }
    
    if bStart >= len(b) {
        return a[ aStart + k - 1]
    }
    
    if k == 1 {
        if a[aStart] < b[bStart] {
            return a[aStart]
        } else {
            return b[bStart]
        }
    }
    
    kHalf := k / 2
    aMid, bMid := math.MaxInt, math.MaxInt
    if aStart + kHalf - 1 < len(a) {
        aMid = a[aStart + kHalf - 1 ]
    }
    
    if bStart + kHalf - 1 < len(b) {
        bMid = b[bStart + kHalf - 1]
    }
    
    if aMid < bMid {
        return getKth( a, b, aStart + kHalf, bStart, k - kHalf)
    } else {
        return getKth( a, b, aStart, bStart + kHalf, k - kHalf)
    }
}