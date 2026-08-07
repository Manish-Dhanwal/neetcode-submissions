func maxArea(heights []int) int {
	left := 0
	right := len(heights)-1
	max := 0
	for (left<right){
		// max := 0
		width := right -left
		area := width * maxmin(heights[left], heights[right])
		if area > max{
			max = area
		}
		if heights[left] < heights[right]{
			left++
		}else{
			right--
		}
	}
	return max
}


func maxmin(x int, y int)int{
	if x<y{
		return x
	}else{
		return y
	}
}