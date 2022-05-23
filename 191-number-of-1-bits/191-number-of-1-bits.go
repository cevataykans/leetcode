func hammingWeight(num uint32) int {
    
    numOfBits := 0
    for num > 0 {
        
        num = num & (num - 1)
        numOfBits++
    }
    return numOfBits
}