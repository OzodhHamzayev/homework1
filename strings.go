package main

import (
	"fmt"
	"strings"
)

func Task3() { 

	//TODO HasPrefix
	//! necha harf tekshirilishini aniqlaydi / uning faqat BOSHI qaraladi, oxiri umuman qaralmaydi
	//!a -> tekshiriluvchi b -tekshiruvchi
	a := "gOlang"
	b := "go"
	resultHasPrefix := strings.HasPrefix(a,b)
	fmt.Println("HasPrefix : ",resultHasPrefix)
	


	//TODO Contains
	//! "a" ning xar qanday indexida "b" ning biror bir xarfi togri kelsa true chiqadi
	//!a -> tekshiriluvchi b -tekshiruvchi
	resultContains := strings.Contains(a,b)
	fmt.Println("Contains : ",resultContains)
	


	//TODO HasSuffix
	//! necha harf tekshirilishini aniqlaydi / uning faqat Oxiri qaraladi, boshi umuman qaralmaydi
	//!a -> tekshiriluvchi b -tekshiruvchi
	resultHasSuffix := strings.HasSuffix(a,b)
	fmt.Println("HasSuffix : ", resultHasSuffix)



	
	//TODO ToUpper
	//! bitta strindagi textlarni xammasini katta xarflarda change qilib beradi
	resultToUpper := strings.ToUpper(a)
	fmt.Println("ToUpper : ",resultToUpper)


	

	//TODO ToLower
	//! bitta strindagi textlarni xammasini kichkina xarflarda change qilib beradi
	resultToLower := strings.ToLower(a)
	fmt.Println("ToLower : ",resultToLower)
	



	cars := []string{"bmw","mazda","volvo","chevrolet"}
	resultJoin := strings.Join(cars, "|")
	fmt.Println("strings.Join :",resultJoin)
	//TODO Join
	//! Strings.Join array ni stringa aylantirib beradi va xar bir indexni orasida qanday qiymat qoshib boladi



	//TODO Split
	//! Strings.Split stringni arrayga aylantiradi va 2-qiymati yozilgan elementni olib tashlaydi
	resultSplit := strings.Split(resultJoin, "|")
	fmt.Println("Split : ",resultSplit)
	

	//TODO TrimSpace
	//! Strings.TrimSpace boshi va oxiridagi bo'shliqlarni olib tashlaydi
	space := "                Hello                      " 
	resultTrimSpace := strings.TrimSpace(space)
	fmt.Println("TrimeSpace : ",resultTrimSpace)




 // You know that len("café") = 5 but len([]rune("café")) = 4.
 // Write a function countCharsAndBytes(s string) (chars, bytes int) that returns both counts. 
 // Test with: "Hello", "Камрон", "Go🚀Код", "café". For each, print whether chars == bytes and explain why/why not.


	names := []string{"Hello", "Камрон", "Go🚀Код", "café"}
	for i := 0; i < len(names); i++ {
		chars, bytes := countCharsAndBytes(names[i])
		fmt.Printf("word  : %s, chars: %d, bytes: %d\n", names[i], chars, bytes)	
	}
}

func countCharsAndBytes(s string) (chars int, bytes int) {
	for i := 0; i < len(s); i++ {
		chars = len([]rune(s))
		bytes = len(s)

	}
	return chars, bytes
}