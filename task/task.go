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


func search(nums []int, target int) int {
	result := -1
	low,high := 0, len(nums)-1
	for (low <= high) { 
		mid := (low+high)/2
		if nums[mid] == target {
			result = mid
			return result
		} else if nums[mid] < target {
			low = mid+1
		} else if nums[mid] > target { 
			high = mid-1
		}
	}
		return result
}

func search2(nums []int, target int) int {
	result := -1
    for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			result = i
			return result
		}
	}
	return result
}

func majorityElement(nums []int) int {
    result := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] >= result {
			result = nums[i]
		}
	}
	return result
}


func fundMin(nums []int) int {
    high,low := len(nums)-1, 0
	result := nums[0]
	for (low <=	high) {
		mid := (high+low)/2
		if nums[low] >= nums[high] {
			low = mid+1 
			if nums[mid] < result {
				result = nums[mid]
			}
		} else if nums[low] < nums[high] {
			high = mid-1 
			if nums[mid] <= result {
				result = nums[mid]
			}
		}
		
	}
	return result
}



func main(){

	nums := []int{4,5,6,7,0,1,2}
	fundMin := fundMin(nums)
	fmt.Println(fundMin)

	// 	nums := []int{-1,0,3,5,9,12}
	// target := 9
	// search := search(nums, target)
	// majorityElement := majorityElement(nums)
	// fmt.Println(search)
	// fmt.Println(majorityElement)
	












	// nums := []int{5,7,7,8,8,8,10}
	// target := 8
	// searchRange := searchRange(nums, target)
	// searchRange2 := searchRange2(nums, target)
	// fmt.Println(searchRange)
	// fmt.Println(searchRange2)
	
}





