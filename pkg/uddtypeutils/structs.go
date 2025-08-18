package uddtypeutils

import "fmt"


/*
	struct declarations:
	**********************************
	structs are user defined datatypes
	that can help in storing bunch of
	stuff together in a proper manner
	(idk how to properly describe)
	in go, we declare a struct using 
	type keyword, followed by the name
	of the datatype.
	we can also write functions that
	can access unexported data ,
	ie the lower case variables
	**********************************
*/
type gasEngine struct {
	kpl uint8
	litres uint8
	ownerInfo owner 
}


/*
	**********************************
	we can call other structs as vars
	inside another struct, allowing 
	for communication between structs?
	**********************************
*/
type owner struct {
	name string
}


func WorkOnStructs () {
	var myEngine gasEngine
	fmt.Printf("k-p-l of %v is %v and its having %v litres\n", myEngine, myEngine.kpl, myEngine.litres)


/*
	variable initialization:
	***********************************************
	there are many ways to initialize a struct var
	1. use the struct format <name>{types...}
	2. use 1 but dont specify types
	3. directly assign values to the struct members
	***********************************************
*/
	var myFullEngine gasEngine = gasEngine{kpl: 12, litres: 128}
	fmt.Printf("k-p-l of %v is %v and its having %v litres\n", myFullEngine, myFullEngine.kpl, myFullEngine.litres)
	myFullEngine.kpl = 20
	fmt.Printf("k-p-l of %v is %v and its having %v litres\n", myFullEngine, myFullEngine.kpl, myFullEngine.litres)

	myFullEngine.ownerInfo.name = "Man"
	fmt.Printf("k-p-l of %v is %v and its having %v litres\n", myFullEngine, myFullEngine.kpl, myFullEngine.litres)

}

/*
the virtual class look-alike in go
is the interface system
in this, we define a new type
which is an interface, and we can
write functions with their signnatures

now if we write a new struct which has
those functions, then its type can be
replaced with that of the interface
*/