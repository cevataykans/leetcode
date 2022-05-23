func hammingWeight(num uint32) int {
    
    numOfBits := 0
    for i := 0; i < 32; i++ {
        
        if num & 1 == 1 {
            
            numOfBits++
        }
        num = num >> 1
    }
    return numOfBits
}