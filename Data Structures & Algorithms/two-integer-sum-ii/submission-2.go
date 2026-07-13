func twoSum(numbers []int, target int) []int {
	seen := make(map[int]int)
	
	// brute force
	// for left < right{
	// 	if numbers[left] + numbers[right] == target{
	// 		return []int{left + 1, right + 1}
	// 	}
	// 	if numbers[left] + numbers[right]>target{
	// 		right--
	// 	}else if numbers[left] + numbers[right]<target{
	// 		left++
	// 	}else{
	// 		left++
	// 		right--
	// 	}
		
	// }
	for i, num := range numbers{
		checkVal := target - num
		value, ok := seen[checkVal];
		 if ok{
			return []int{value + 1, i + 1,}
		  }
		  seen[num] = i
	}
	return []int{}
}
