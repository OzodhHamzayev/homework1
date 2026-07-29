package main

import (
	"errors"
	"fmt"
	"strings"
)

func firstFunction() error { 
	return fmt.Errorf("something went wrong")
}

func secondFunction() error {
	firstFunction := firstFunction()
	if firstFunction != nil {
		return fmt.Errorf("something went wrong: %w", firstFunction)
	}
	return nil
}



type NoText struct {
	Message string
}

type TooBig struct { 
	Size int
}

type error interface { 
	Error() string
}

func (n *NoText) Error() string  {
	return n.Message
}

func (s *TooBig) Error() string {
	return fmt.Sprintf("juda katta: %d", s.Size)
}


func Check(name string, size int) (string, int, error) {
	if !strings.HasSuffix(name, "/txt") {
		return "", 0 , &NoText{Message: "ends with .txt"}
	} else if size < 1000 {
		return "", 0, &TooBig{Size: size}
	}
	return name, size, nil
}








func main() {

	var s  *TooBig
	var nt *NoText
	name, size, err := Check("go.txt", 1001)
	if errors.As(err, &s) {
		fmt.Println(err)
	}else if errors.As(err, &nt) {
		fmt.Println(err)
	}else {
		fmt.Println(name, size)
	}
	fmt.Println(err)











	// fmt.Println(firstFunction())
	// fmt.Println(secondFunction())

	// message, err := loadFile("go.mod")
	// fmt.Println("message:", message)
	// fmt.Println("error:", err)


	// message1, err1 := AlotFiles("go.txt")
	// if errors.Is(err1, ErrNoText) {
	// 	fmt.Println("asda")
	// }else if  errors.Is(err1, ErrNoAccess ) {
	// 	fmt.Println(ErrNoAccess)
	// } else 	{
	// 	fmt.Println(message1)
	// }

	




	
	// name, age , err2 := Describe("Ozod", -1)
	// fmt.Println("name:", name)
	// fmt.Println("age:", age)
	// fmt.Println("error:", err2)














}




















func loadFile(name string) (string, error) {
	if !strings.HasSuffix(name, ".text") {
			 return "", &NoText{ Message:"Opening not txt files not allowed" }
	}
	
	return name, nil
}
var ErrNoText = errors.New("fayl matn emas")
var ErrTooBig = errors.New("fayl juda katta")
var ErrNoAccess = errors.New("ruxsat yo'q")

func AlotFiles(name string) (string, error) { 
	if !strings.HasSuffix(name, ".txt") {
		return "", ErrNoText
	} else if len(name) > 100 {
		return "", ErrTooBig
	} else if name == "go.txt" {
		return "", ErrNoAccess
	}

	return name, nil
}

func Describe(name string, age int) (string, int, error) {
	if name == "" || age < 0 {
		return "", 0, &NoText{Message: "please fill the form correctly"}
	}
	return name, age, nil
}





