package dtypeutils
import "fmt"

func WorkOnMaps() (map[string]uint8){
	/* 
	map declarations
	can be instantiated using make
	can also be directly populated with values

	maps, just like in c++, will initialize something with 0
	if the thing does not exist as a key, so we need to be
	careful about that, luckily can be checked using ok boolean

	can delete stuff using delete keyword
	*/
	var Map map[string]uint8 = make(map[string]uint8)
	fmt.Println(Map)

	var Map2 = map[string]uint8{"Adam": 23, "Me": 19}
	fmt.Println(Map2)
	fmt.Println(Map2["Me"])
	var age, ok = Map2["Adam"]
	fmt.Printf("found age: %v, is person present? %v \n", age, ok)
	delete(Map2, "Adam")
	age, ok = Map2["Adam"]
	fmt.Printf("found age: %v, is person present? %v \n", age, ok)
	
	var Map3 = map[string]uint8{"Person1": 23, "Person2": 19, "Person3": 28}
	return Map3
}	