package main

import "fmt"


//! search range -> 0(n)
func searchRange(nums []int, target int) []int {
	// high, low :=  len(nums)-1, 0
	// first, last := -1,-1
	result := []int{-1,-1}
		for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			if result[0] == -1 {
				result[0] = i
			}else{
				result[1] = i
			}
		}
	}

	return result
}



func main(){
	nums := []int{5,7,7,8,8,10}
	target := 8
	searchRange := searchRange(nums, target)
	fmt.Println(searchRange)
}



