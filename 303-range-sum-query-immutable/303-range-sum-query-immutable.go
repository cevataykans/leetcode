type NumArray struct {
    
    sumUntilIndex []int
}


func Constructor(nums []int) NumArray {
    
    sums := make( []int, len( nums) + 1, len( nums) + 1)
    sums[ 1] = nums[ 0]
    for i := 1; i < len( nums); i++ {
        
        sums[ i + 1] += sums[ i] + nums[ i]
    }
    
    newNumArr := NumArray{
        
        sumUntilIndex: sums,
    }
    return newNumArr
}


func (this *NumArray) SumRange(left int, right int) int {
    
    return this.sumUntilIndex[ right + 1] - this.sumUntilIndex[ left]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */