package main
import (
	"github.com/ayush-bothra/go_practice_main/pkg/dtypeutils"
	// "github.com/ayush-bothra/go_practice_main/pkg/cntrlflowutils"
)

/*
*************************************************************************
This is a for practice file only, u can un-comment any part of the code 
(as long as its not causing any issues to the go compiler), run it using
go run <file_path> and see the results, each of the files have some level
of comments explaining details about the workings of code in go, these 
can be referred when required.
each of the functionalities can be checked using CTRL+CLICK
*************************************************************************
*/

func main() {
	/* 
	math stuff to check how go works, error handling
	*****************************************************************************************
	var value1 int = 10
	var value2 int = 5
	var quotient, remainder, err = mathutils.IntegerDivision(value1, value2)

	 if - else ladder in go 
	***************************************************************************************** 
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0{
		fmt.Printf("The quotient is: %v \n", quotient)
	} else {
		fmt.Printf("The quotient and remainder are: %v and %v \n", quotient, remainder)
	} 
	***************************************************************************************** 
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The quotient is: %v \n", quotient)
	default:
		fmt.Printf("The quotient and remainder are: %v and %v \n", quotient, remainder)
	}
	/*****************************************************************************************
	 the above is the switch statement in go, they also have conditional switch statements,
		where the syntax is switch conditional_variable {}, and the cases depend on the val of
		the conditional variable
	******************************************************************************************
	*/

	// dtypeutils.WorkOnArrays()
	// dtypeutils.WorkOnMaps()
	// cntrlflowutils.WorkOnForLoops()
	dtypeutils.WorkOnStrings()
}

