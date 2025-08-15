package cntrlflowutils

import (
	"fmt"
	"github.com/ayush-bothra/go_practice_main/pkg/dtypeutils"
)
/*

files in the same package folder
share the same package namespace
so they have access to other file
variables and functions (in the global scope)

*/

func WorkOnForLoops() {
	var newMap = dtypeutils.WorkOnMaps()
	fmt.Println(newMap)

	/*
	for loop declarations
	**************************************************
	the for loop syntax is somewhat similar to c++
	except for the fact that u can have multiple 
	variables like in python 
	and u can write c++ loop structure BUT without the 
	braces
	***************************************************
	*/

	for name, age := range newMap {
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	/*
	*********************************************************
	go does not employ the usage of while loops, but u can
	bend the for loop a little to turn it into some sort of a
	while loopish like structure, also u can write for loops 
	in go without their conditional structure, instead it is 
	implied that the loop is infinite and inner break 
	statements decide what is to be done of the loop itself
	altho go does have staticcheks which can catch infinite 
	loops
	*********************************************************
	*/

	var i int8 = 0
	fmt.Printf("While loop form: \n")
	for i < 10 {
		fmt.Printf("%v\t",i)
		if ((i + 1) % 10 == 0) {
			fmt.Printf("\n")
		}
		i++
	}

	/* 
	standard declarations 
	**************************************
	first one is pre go 1.21
	next one is paritial closed
	representation of the same but using
	ranges
	**************************************
	*/

	fmt.Printf("Standard for loop: \n")
	for i = int8(0); i < 10; i++ {
		fmt.Printf("%v\t",i)
		if ((i + 1) % 10 == 0) {
			fmt.Printf("\n")
		}
	}

	fmt.Printf("Using range: \n")
	for i = range int8(10) {
		fmt.Printf("%v\t",i)
		if ((i + 1) % 10 == 0) {
			fmt.Printf("\n")
		}
	}
}