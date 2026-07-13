func longestConsecutive(nums []int) int {
// brute force
longest := 0

// for _, num := range nums{
// 	current := num
// 	lenght := 1

// 	 for contain(nums, current+1){
// 		current++
// 		lenght++
// 	 }

// 	if longest < lenght{
// 		longest = lenght
// 	}
// }

// return longest

set := make(map[int]bool)
for _, num := range nums{
	set[num] = true
}
for _, num := range nums{
	if set[num-1]{
		continue
	}
	current := num
	length := 1
	
	for set[current+1]{
		current++
		length++
	}

	if longest < length{
		longest = length
	}

}
return longest
}


// func contain(nums []int, target int)bool{
// 	for _, num := range nums{
// 		if num == target{
// 			return true
// 		}
// 	}
// 	return false
// }