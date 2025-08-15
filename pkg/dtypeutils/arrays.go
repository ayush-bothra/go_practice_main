package dtypeutils
import "fmt"

func WorkOnArrays() {
	/* 
	array declarations 
	each run will print differnt addr space loc
	due to ASLR in Windows
	*******************************************
	var intArr [3]int16 = [3]int16{1, 2, 3}
	fmt.Println(intArr)
	fmt.Println(&intArr[0]);
	fmt.Println(&intArr[1]);
	*******************************************
	*/

	/* 
	slices declarations
	this is a wrapper around the array
	this helps in dynamic allocation like c++ vectors
	*/
	var intSlice []int16 = []int16{1, 2, 3}
	fmt.Printf("the length of the Slice is %v and capacity is %v \n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 4)
	fmt.Println(intSlice);
	intSlice = append(intSlice, 5)
	fmt.Println(intSlice)
	fmt.Printf("the length of the Slice is %v and capacity is %v \n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, intSlice...)
	fmt.Println(intSlice)

	// make takes in the dtyp, the length and the capacity of the slice
	var intNewSlice []int16 = make([]int16, 3, 8)
	fmt.Printf("Slice: %v,\t length: %v,\t capacity: %v \n", intNewSlice, len(intNewSlice), cap(intNewSlice))
}