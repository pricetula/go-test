package main

import "fmt"

type exampleType struct {
	name string
	age  int
}

func main() {
	// Variable map1 has a map of key-value pairs with the key being a string type and its value being an int type
	map1 := make(map[string]int)
	map1["mango"] = 90
	map1["orange"] = 291
	fmt.Println("map1 = ", map1)
	fmt.Println("map1[\"mango\"] = ", map1["mango"])
	fmt.Println("len(map1) = ", len(map1))
	// Deleted an item with key mango
	delete(map1, "mango")
	//  item with key mango shouldn't exist
	fmt.Println("map1 = ", map1)
	// item with key mango was deleted hence referncing value on this key will be the initial value
	fmt.Println("map1[\"mango\"] = ", map1["mango"])
	// length should not consider item with mango key
	fmt.Println("len(map1) = ", len(map1))

	var map2 map[int]exampleType = map[int]exampleType{
		1: exampleType{name: "mango", age: 20},
		2: exampleType{name: "apple", age: 34},
	}
	fmt.Println("map2 = ", map2)
	fmt.Println("map2[\"mango\"] = ", map2[1])
	fmt.Println("len(map2) = ", len(map2))
	value, exists := map2[1]
	fmt.Println("value, exists := map2[1]; value = ", value, " Exists = ", exists)
}
