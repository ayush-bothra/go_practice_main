package dtypeutils

import "fmt"

func WorkOnPointers () {
	/*
	pointers in go work the same as
	in c++ or c, so no need to write
	extra stuff :)
	*/

	// do this to avoid nil pointer exceps
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value of p is : %v and its pointing to %v\n", p, *p)
	fmt.Printf("The value of i is %v\n", i)

	*p = 10
	fmt.Printf("The value of p is : %v and its pointing to %v\n", p, *p)

	p = &i
	*p += 10
	fmt.Printf("The value of i is %v\n", i)

	/*
	important note, sliceCopies use pointers to create shallow copies ? 
	so changing one value leads to changes in the other
	*/

	/*
	usual pass by reference works here as well, however
	instead of having the liberty of using & in func
	signature (looking at u c/c++) we just return a pointer
	this makes it much simpler to remember.
	*/
}