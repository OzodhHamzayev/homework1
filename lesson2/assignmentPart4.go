package main

import (
	"errors"
	"fmt"
)

func Task4() { 
	//DNS-inspired function: Write validateTTL(ttl int) (string, error) that returns:
	//error if ttl < 0
	//"very short" if 0 ≤ ttl < 60
	//"short" if 60 ≤ ttl < 300
	//"normal" if 300 ≤ ttl < 3600
	//"long" if 3600 ≤ ttl < 86400
	//"very long" if ttl ≥ 86400
	//Test with: -1, 30, 60, 300, 3600, 86400.

	nums := []int{-1, 0, 30, 60, 300, 3600, 86400.}
	for i := 0; i < len(nums); i++ {
		result, error := validateTTL(nums[i])
		if error != nil {
			fmt.Println(error)
		} 
		fmt.Println(result)
	}

}

func validateTTL(ttl int) (string, error) {
	if ttl < 0 {
		return "", errors.New("0 dan kichik")
	}
	if  0 <= ttl && ttl < 60 {
		return "very short", nil
	} else if 60 <= ttl && ttl < 300 {
		return "short", nil
	} else if 300 <= ttl && ttl < 3600 {
		return "normal", nil
	} else if 3600 <= ttl && ttl < 86400 {
		return "very long", nil
	}
	return "", nil
}