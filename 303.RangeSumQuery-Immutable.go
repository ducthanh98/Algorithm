package main

type NumArray struct {
	mapNumber map[int64]int
}


func Constructor(nums []int) NumArray {
	arr := make(map[int]int,len(nums))

	for i:=0;i<len(nums);i++ {
		arr[i] = nums[i]
	}

	return NumArray{
		mapNumber: arr
	}   
}


func (this *NumArray) SumRange(left int, right int) int {
	result := 0
    for i:= left;i<=right;i++ {
		result += this.mapNumber[i]
	}
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

func main(){

}