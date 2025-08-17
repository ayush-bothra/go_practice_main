package dtypeutils

import (
	"fmt"
	"strings"
)

func WorkOnStrings() {
	/*
	********************************************
	string declarations
	strings in go are encoded in the UTF-8
	encoding, therefore we can also make use 
	of special characters that ASCII does not
	let us make use of.
	UTF-8 is a variable type encoding that 
	stores the size of the encoding in the 2
	MSBs of each byte. This is why alot of the
	indexes will get skipped when making use of
	the range keyword while looping over a 
	string utilizing the special characters.
	********************************************
	*/
	var newString = "ṇinṅh";
	fmt.Printf("Normal String: %v \n", newString)
	var indexed = newString[1]
	fmt.Printf("%v, %T \n", indexed, indexed)

	for i, v := range newString {
		fmt.Println(i, v)
	}
	fmt.Printf("length of Normal string: %v\n", len(newString))

	/*
	******************************************
	To fix the indexing issues that range has
	while looping over a string, we convert it
	into a set (array really) of runes (uint8)
	******************************************
	*/

	var runeString = []rune(newString)
	fmt.Printf("Rune String : %v \n", runeString)
	var runeindexed = runeString[1]
	fmt.Printf("%v, %T \n", runeindexed, runeindexed)

	for i, v := range runeString {
		fmt.Println(i, v)
	}
	fmt.Printf("length of Rune string: %v\n", len(runeString))

	/*
	***************************************************
	next is string building: can be done by declaring
	an array of char (can be called rune) values with
	a variable name.
	strings can be concatenated using the '+' symbol 
	(no overloading btw, go doesnt support overloading)

	NOTE: strings are IMMUTABLE (pls remember this)
	***************************************************
	*/

	var strSlice = []string{"s", "u", "b", "s", "t", "r"}
	fmt.Printf("This is a string built using slices: %v\n", strSlice)
	var catSlice = "This is a "
	// catSlice += strSlice is not allowed for slices

	for i := range strSlice {
		catSlice += strSlice[i]
	}
	fmt.Println(catSlice)

	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	// This is faster than normal loop appending
	var catSlice2 = strBuilder.String()
	fmt.Println(catSlice2)
}