package main

import (
	"fmt"
)

//! Binary search
func searchInsert(nums []int, target int) int {
    high := len(nums)-1 // 4
	low := 0 // 0
	for(low<=high){
		mid := (high+low)/2 // 2
		if nums[mid] == target {
			return mid
		}else if nums[mid] < target {
			low = mid+1
		} else if nums[mid] > target {
			high = mid-1
		}
	}
	return low
}









//!single Number -> not fulll
func singleNumber(nums []int) int {
	a := nums[0]
    for i := 1; i <= len(nums)-1; i++ {
		if a == nums[i] {
			a += 1
		}else if a != nums[i]{
			a = nums[i]
			break
		}
	}
	return a 
}

//! not fulll
// func singleNumbers(nums []int) int {
// 	b := 0
//     for i := 0; i <= len(nums)-1; i++ {
// 		for k := 1; k < len(nums)-1; k++ {
// 			if nums[i] == nums[k] {
// 				nums[i] = 0
// 				nums[k] = 0
// 			} else if nums[i] != nums[k]{
// 				b = nums[i]
// 			}
// 		}
// 	}
// 	return a 
// }




//! insert Elements
func InsertElements(nums []int, target int) []int { 
	nums = append(nums, 0)
	r := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] <= target && nums[i+1] > target{
			r = nums[i+1]
			fmt.Println(r)
			for k := len(nums)-1; k >= i+1; k-- {
				nums[k]=nums[k-1]
				if r == nums[k] {
					target = k
				}
			}
				break
		}
	}
	return nums
}


func main() { 

	nums := []int{1,2,3,4,5,6,7,8,9}
	target := 8
	InsertElements := InsertElements(nums, target)
	fmt.Println(InsertElements)


}







































	// nums := []int{1,2,3,4,5,6}
	// target := 6
	// searchInsert := searchInsert(nums,target)
	// if searchInsert == -1 {
	// 	fmt.Println("number not found")
	// }else {
	// 	fmt.Println(nums[searchInsert])
	// 	fmt.Println(searchInsert)
	// }







	// nums := []int{4,1,2,1,2}
	// singleNumbers := singleNumbers(nums)
	// singleNumber := singleNumber(nums)
	// fmt.Println(singleNumber)
	// fmt.Println(singleNumbers)



// ! binary search
// func BinarySearch(arr []int, x int) int {
// 	high := len(arr)-1
// 	low := 0
// 	mid := (high+low)/2
// 	if arr[high] < x || arr[low] >= x{
// 		return -1
// 	} else {
// 		for low <= high {
// 		if x == arr[mid]{
// 			return mid
// 		} else if x > arr[mid] {
// 			low = mid+1
// 		} else if x < arr[mid] {
// 			high = mid-1
// 		}
// 	}
// 	return -1
// 	}

// }


// //! PlusOne
// func plusOne(digits []int) []int {
// 	for i := len(digits)-1; i >= 0; i-- {
// 		if digits[i] < 9 {
// 			digits[i] += 1
// 			break
// 		}else if digits[i] == 9 {
// 			digits[i] = 0
// 			if digits[0] == 0 {
// 				digits[0] = 1 
// 				digits = append(digits, 0)
// 			}
// 		}
// 	}
// 	return digits
// }










// 	!plusOne
// 	digits := []int{9,9,3,2,9}
// 	plusOne := plusOne(digits)
// 	fmt.Println(plusOne)


// ! twoSum
// func twoSum(nums []int, target int) []int {
//     result := []int{}
//     for i := 0; i < len(nums); i++ {
//         for k := i+1; k <= len(nums)-1; k++ { 
//             if nums[i]+nums[k] == target {
//                 result = append(result, i,k)
//                 return result
//             }
//         }
//     }
//     return result
// }


// //!Remove Duplicates
// func removeDuplicates(nums []int) []int {
// 	nums = append(nums, 0)
// 	k := 0
// 	uniqueElements := []int{}
// 	newElements := []int{}
// 	for i := 0; i < len(nums)-1; i++ {
// 		if nums[i] == nums[i+1] {
// 			newElements = append(newElements, nums[i])
// 		}else if nums[i] != nums[i+1] {
// 			k++
// 			uniqueElements = append(uniqueElements, nums[i])
// 		}

// 	}
// 	fmt.Println(newElements)
// 	for i := 0; i < len(newElements)-1; i++ {
// 			uniqueElements = append(uniqueElements, newElements[i])
// 	}


// 	return uniqueElements
// }




// 	Task1() 
// 	Task2()
// 	Task3()
// 	Task4()
// 	Task5()		




// 	numbers := []int{0,0,1,1,1,2,2,3,3,4}
// 		result := removeDuplicates(numbers)
// 		fmt.Println(result)


// 	numbers := []int{3,2,4}
// 	target := 6
// 	result := twoSum(numbers, target)
// 	fmt.Println(result)





// b := 1
// v := 3
// result1 := []int{}
// result1 = append(result1, b,v)
// fmt.Println(result1)
// //!binary search

// 	numbers := []int{1,2,3,4,5,6,7,8,9}
// 	num := -50

// 	result := BinarySearch(numbers, num)
// 	if result != -1 {
// 		fmt.Println("number is found:", result)
// 	} else {
// 		fmt.Println("number not found")
// 	}


// }



// ! sorting
// 	numbers := []int{10,29,19,100,9,8,7,6,5,39,4,100,120,3,2,1}

// 	fmt.Print(numbers, "\n")

// 	for j := 0; j < len(numbers)-1; j++ {	
// 		for i := 0; i < len(numbers)-1; i++ { 
// 			if numbers[i] > numbers[i+1] {
// 				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
// 			}			
// 		}
// 	}
	
// 	fmt.Println(numbers)







// 	!majority elements
// 	equalNumbers := []int{1,1,1,1,2,2,3,3,4,4,4}

// 	winners := []int{}
// 	result := 1
// 	p := &result
// 	equalNumbers = append(equalNumbers, 0)
// 	for i := 0; i < len(equalNumbers)-1; i++ { 
// 		if equalNumbers[i] == equalNumbers[i+1] {
// 			(*p)++		
// 		}	else if  equalNumbers[i] != equalNumbers[i+1] {
// 			winners = append(winners, *p)
// 			*p = 1
// 		}
// 	}
// 	fmt.Println(winners)
// 	lastResult := 0
// 	for i := 0; i < len(winners); i++ {
// 			if winners[i] > lastResult {
// 				lastResult = winners[i]
// 			}
// 	}
// 	fmt.Println(lastResult)





// ! majority elements 2
// 	majorityNum := []int{1,1,1,1,2,2,2,3,3}

// 	result := 0

// 	for i := 0; i < len(majorityNum); i++ {
// 		for j := 0; j < len(majorityNum); j++ {
// 			if majorityNum[i] == 1 {
// 				result++
// 			}
// 		}
// 	}

// 				fmt.Println(result)


// 	const name = 12
// 	fmt.Println(name)









// !majority elements 3

// 	differentNumber := []int{0,1,2,2,3,4,5,6,6,7,7,8,9,10,10}

// 	for i := 0; i < len(differentNumber); i++ { 
// 		for j := i+1; j < len(differentNumber); j++ {
// 			if differentNumber[i] == differentNumber[j] {	
// 				differentNumber[i] = 0
// 				differentNumber[j] = 0
// 			}
// 		}
// 	}
// 	fmt.Println(differentNumber)

// 	newArray := []int{}
// 	for i := 0; i < len(differentNumber); i++ {
// 		if differentNumber[i] != 0 {
// 			newArray = append(newArray, differentNumber[i])
// 		}
// 	}
// 	fmt.Println(newArray)
















// 	!deletion 0(n)

// 	deletion := []int{1,2,3,4,5,6,7,8,9}
// 	deletionIndex := 3

// 	for i := deletionIndex; i < len(deletion)-1; i++ { 
// 		deletion[i] = deletion[i+1]
// 	}
// 	deletion = deletion[:len(deletion)-1]
// 	fmt.Println(deletion)


// 	number1 := []int{1,2,3,4,5,6,7,8}

// 	n := number1[  :7-1]
// 	fmt.Println(n)

// 	! delete from the end -> 0(1)

// 	endingNumber := []int{1,2,3,4,5,6,7,8,9}

// 	fmt.Println(endingNumber)

// 	length := len(endingNumber)
// 	endingNumber = endingNumber[:length-1]
// 	fmt.Println(endingNumber)







// 	! insertion  ->
// 	!  2 xil usul. ketma ket qilish yani 1 ta elementni insert qilgandan keyin qolganlari faqat bittadan
// 	!siljishi kerak keyingi indexga yoki 0-indexni olib keyingi eng oxirgi index ga qoyish kerak

// 	! 1
// 	nums := []int{11,23,34,41,15,16,75,58,19}
// 	newElement := 0
// 	nums = append(nums, newElement)
// 	resultN := 5
// 	fmt.Println(nums)
// 	for i := len(nums)-1; i > 0; i-- { 
// 		nums[i] = nums[i-1]
// 	}
// 	nums[0] = resultN
// 		fmt.Println(nums)



// 	! 2	
// 	nums1 := []int{11,23,34,41,15,16,75,58,19}
// 	newNumber := 0
// 	nums1 = append(nums1, newNumber)
// 	fmt.Println(nums1)
// 	result := len(nums1)-1
// 	nums1[result] = nums1[0]
// 	nums1[0] = 5
// 	fmt.Println(nums1)



// 	! 2-1
// 	nums2 := []int{11,23,34,41,15,16,75,58,19} 
// 	newNumber1 := 0
// 	newNum := 3333
// 	nums2 = append(nums2, newNumber1)
// 	result2 := len(nums2)-1
// 	for i := 0; i < len(nums2)-1; i++ { 
// 		if nums2[i] == 34 {
// 			nums2[result2] = nums2[i]
// 			nums2[i] = newNum
// 			break
// 		}
// 	}
// 	fmt.Println(nums2)

// 	!2-2
	
// 	nums3 := []int{11,23,34,41,,115,16,75,589} 
// 	newNum3 := 3333
// 	insertIndex := 3
// 	nums3 = append(nums3, 0)
// 	for i := len(nums3)-1; i > insertIndex; i-- { 
// 			nums3[i] = nums3[i-1]
// 	}
// 		nums3[insertIndex] = newNum3

// 	fmt.Println(nums3)









// 			!traversal
// 		traversal := []int{1,2,3,4,5,6,7,8,9}

// 		for i := 0; i < len(traversal); i++ { 
// 			fmt.Println(traversal[i])
// 		}



// 	!find the largest element an array	

// 	elements := []int{10,20,312,12,41,901,999,123,312,421,-10}
// 	laregestElement := elements[0]
// 	for i := 0; i < len(elements); i++ {
// 		if elements[i] >= laregestElement  {
// 			laregestElement = elements[i]
// 		}
// 	}
// 	fmt.Println(laregestElement)

// 	smallerElement := elements[0]
// 	for i := 0; i < len(elements); i++ {
// 		if elements[i] <= smallerElement {
// 			smallerElement = elements[i]
// 		}
// 	}
// 	fmt.Println(smallerElement)




// 	!Reverse the elements of an array

// 	sortedElements := []int{1,2,3,4,5,6,7}
// 	reverseElements := []int{}
// 	for i := len(sortedElements)-1; i >= 0; i-- {
// 		reverseElements = append(reverseElements, sortedElements[i])
// 	}
// 	fmt.Println(reverseElements)





// 	!Search  for a specific element in  an array
// 	searchingElements := []int{1,2,3,4,5,6,7,8,9}
// 	newNumber := -6
// 	found := false
// 	for i := 0; i < len(searchingElements); i++ {
// 		if searchingElements[i] == newNumber {
// 			found = true
// 			break
// 		}
// 	}

// 	if found == true {
// 		fmt.Printf("%d was found\n", newNumber)
// 	}else{
// 		fmt.Printf("%d was not found in the array\n", newNumber)
// 	}
// 	! Calculate the sum of all elements in an array.

// 	numbers := []int{1,2,3,12,4,2,4,2,12,12,35,97,62}
// 	sum := 0
// 	for i := 0; i < len(numbers); i++ { 
// 		sum += numbers[i]
// 	}
// 	fmt.Println(sum)

	



//  	newSlice := []int{1,2,3,4,5,6}
// 		sum := []int{}
// 	for i := 0; i < len(newSlice); i++ {
// 		for j := i+1; j < len(newSlice); j++ {
// 			sum = append(sum, newSlice[i], newSlice[j])
			
// 		}
// 	}
// 		fmt.Printf("sum: %d\n\n", sum)





// 		searching := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
// 		resultNumber := 15

// 		for i := len(searching)-1; i > 0; i=i/2 { 
// 				if searching[i]/2 >= resultNumber {
						
// 				}else if searching[i] <= resultNumber {
// 					fmt.Println(searching[i])

// 				}
// 		}




// 		!two-dimensional

// 		matrix := [2][3]int{
// 			{1,2,3},
// 			{4,5,6},
// 		}
// 		fmt.Println(matrix[1][0])


// 		!pointer with array

// 		addressOfNumber := []int{0,1,2,3,4,5,6,7,8,9}
// 		a := &addressOfNumber[2]
//  		fmt.Println(addressOfNumber)
// 		fmt.Println(a)
// 		fmt.Println(*a)



// 		!deletion

// 		outdatedCars := [6]string{"BMW", "CHevrolet", "Mazda", "Mustang", "KIA", "Ferrari"}
// 		removeCars := []string{}
// 		keptCars := []string{}
// 		for _, v := range outdatedCars {
// 			if len(v) <= 4 {
// 				removeCars = append(removeCars, v)
// 			}else {
// 				keptCars = append(keptCars, v)
// 			}
// 		}
		


// 			fmt.Println("new car:", removeCars)

// 			fmt.Println("removed  cars:", keptCars)



// 		outdatedCars1 := [6]string{"BMW", "CHevrolet", "Mazda", "Mustang", "KIA", "Ferrari"}
// 		removeCars1 := []string{}
// 		for i, v := range outdatedCars1 {
// 			if len(v) <= 4 {
// 				removeCars1 = append(removeCars1, v)
// 				outdatedCars1[i] = "false"
// 			}
// 		}
// 		writeIndex := 0
// 		for i := 0; i < len(outdatedCars1); i++ {
//     		if outdatedCars1[i] != "false" {
//         		outdatedCars1[writeIndex] = outdatedCars1[i]
//         		writeIndex++

//     		}
// 		}	
	
		
// 			for writeIndex < len(outdatedCars1) {

// 				outdatedCars1[writeIndex] = ""
// 				writeIndex++
// 			}

// 		fmt.Println("Siljitilgandan keyin:", outdatedCars1)







// 		!traversal
// 		number := []int{1,2,3,4,5,6,7,8,9}

// 		for i := 0; i < len(number); i++ {
// 			fmt.Println(number[i])
// 		}



// 		!insertion
// 		name := []int{10, 20, 30, 40, 50}
		
// 		name = append(name, 0)
// 		for i := len(name)-1; i > 0; i--{
// 			name[i] = name[i-1]

// 			}
// 			name[0] = 5
// 		fmt.Println(name)


// 		name1 := []int{10, 20, 30, 40, 50}
// 		fmt.Println(name1[4])
// 		fmt.Println(len(name1))

// 		for i := 0; i < len(name1); i++ {
// 			fmt.Println(i)
// 			fmt.Println(name1[i])
// 		} 


// 		for i := len(name1)-1; i > 0; i-- {
// 			// fmt.Println(i)
// 			fmt.Println(name1[i])
// 		}


		
		
// 	}


	

// 		a := make([]int, 5)

// 			n := len(a)
// 			result := []int{}
// 			for i := 0; i < n; i++ {
// 				a[i]++

// 				result = append(result, a[i])
// 				fmt.Println(result)
// 				}
			

// 		changeName := &name
// 		for _, v := range changeName {
// 			fmt.Println(v)


// 	changeName[0] = 19
// 		fmt.Println(changeName)
// 		fmt.Println(name)




// 		name1 := make(map[string]string)

// 		name1["name"]="i5"
// 		name1["brand"]="BMW"
// 		name1["age"]="409"
// 		name1["reputation"]="good"

// 		for _, v := range name1 {
// 			fmt.Println(v)
// 		}
		
// 		full := map[string][]string {
// 			"brand" : []string{"chevrolet", "matiz"},
// 			"age" : []string{"42", "32"},
// 			"reputition" : []string{"excellent"},
// 		}

// 		for _, v := range full {
// 			fmt.Println(v)

// 		}
// 	}