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


func searchRange2(nums []int, target int) []int {
	low, high := 0,len(nums)-1
	result := []int{-1,-1}
	for (low <= high) {
		mid := (low+high)/2
		if nums[mid] == target  {
			high = mid-1
			if result[0] == -1 {
				result[0] = mid		
			}else {
				result[1] = mid
			}
			fmt.Println(result)
		} else if nums[mid] < target {
			low = mid+1
		} else if nums[mid] > target { 
			high = mid-1
		}
	}
	return result
}


func main(){
	nums := []int{5,7,7,8,8,8,10}
	target := 8
	searchRange := searchRange(nums, target)
	searchRange2 := searchRange2(nums, target)
	fmt.Println(searchRange)
	fmt.Println(searchRange2)
}



