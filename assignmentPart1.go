package main

// import (
// 	"fmt"
// )



// func main( ) { 

// //! 1
// //!  Declare a variable of type byte (= uint8). Assign it the ASCII code of 'A' (= 65). 
// //! Convert it to string — what do you get? Convert string("Z"[0]) to int — what do you get? 
// //! Explain what you're seeing.

// 	var a uint8 = 65
// 	b := string(a)
// 	fmt.Println(b)
// //TODO result : 65 ni stringa aylantirsa A chiqadi




// //! 2 
// 	//! Write code that converts between these types and prints all intermediate values: int → float64 → int → int8. 
// 	//! Use the value 300. What happens at the last step? Why?


// 	var num int = 300
// 	floatNum := float64(num)
// 	fmt.Println(floatNum)
// 	intNum := int(floatNum)
// 	fmt.Println(intNum)
// 	endingNum := int8(intNum)
// 	fmt.Println(endingNum)
// //TODO the result is 44 because int8 only includes from -128 to 127











// //! 3
// //!Declare: a bool for whether a patient is active, an int patient ID, a float64 for temperature (36.6),
// //!  a string for name. Print all using a single fmt.Printf call with multiple format verbs.
// 	var condition bool 
// 	id := 1
// 	temperatureOfPatient :=36.6
// 	name := "Ozod"


// 	fmt.Printf(
// 		"Patient active: %t, Patient ID: %d, Temperature: %.1f, Name: %s\n", condition,id,temperatureOfPatient,name)
// //TODO

// 	}