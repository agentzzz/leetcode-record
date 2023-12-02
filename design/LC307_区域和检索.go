package design

type NumArray struct {
	Nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		Nums: nums,
	}
}

func (this *NumArray) Update(index int, val int) {
	this.Nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	res := 0
	for i := left; i <= right; i++ {
		res += this.Nums[i]
	}
	return res
}
