package main

// import (
// 	"errors"
// 	"fmt"
// )


// func main() {

// //! 4
// //!Write isPrime(n int) bool. No library functions — use a loop. 
// //! Test: 1, 2, 3, 4, 17, 100. Then write countPrimes(max int) int that counts all primes up to max.
// //!TODO
// 	numbers := []int{1, 2, 3, 4, 17, 100, 105}
// 	for i := 0; i < len(numbers); i++ {
// 		a := isPrime(numbers[i])
// 		countPrimes := countPrimes(i)
// 		fmt.Println(a)
// 		fmt.Println(countPrimes)
// 	}

// 	result, error1 := repeat("go", -2)
// 	if error1 != nil{ 
// 		fmt.Println(error1)
// 	} else {
// 		fmt.Println(result)
// 	}

// }

// 	func isPrime(n int) bool { 
// 		if n <= 1  {
// 			return false
// 		}
// 		for i := 2; i < n-1; i++ {
// 			if n % i == 0{
// 				return false
// 			}
// 		}
// 		return true
// 	}
// 	func countPrimes(max int) int { 
// 		count := 0 
// 		for i := 0; i <= max; i++ {
// 			if isPrime(i) == true {
// 				count++
// 			}
// 		}
// 		return count
// 	}
// //! 5
// //! Write a function repeat(s string, n int) (string, error). 
// //! Return error if n < 0. Without using strings.
// //! Repeat — build it manually using a loop (hint: use + to concatenate strings). 
// //! Test: repeat("go", 3), repeat("ha", 0), repeat("x", -1).



// func repeat(s string, n int) (string, error) {
// 	if n < 0 {
// 			return "", errors.New("0 dan katta bolish kerak qiymat uni ko'paytirish uchun")
// 		}
// 	result := ""
// 	for i := 0; i < n; i++ {
// 			result += s
// 	}
// 	return result, nil
// }