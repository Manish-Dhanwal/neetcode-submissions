func productExceptSelf(nums []int) []int {

	prefix := 1
	ans := make([]int, len(nums))

	for i := 0; i < len(nums); i++{
		ans[i] = prefix
		prefix *= nums[i]
	}
	suffix := 1
	for i := len(nums)-1; i >=0; i--{
		ans[i] *= suffix
		suffix *= nums[i]
	} 

	return ans
}
