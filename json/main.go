package main

import "fmt"

type addressDetailType struct {
	street  string
	country string
}

type bioDetailType struct {
	age  int
	race string
}

type hobbyType struct {
	name string
}

type hobbyListType = []hobbyType

type userDetailType struct {
	address addressDetailType
	bio     bioDetailType
	hobbies hobbyListType
}

type userDetailsType = []userDetailType

func main() {
	userDetails := []userDetailsType{}
	for _, value := range [...]int{1, 2, 3, 4} {
		fmt.Println("userD", userDetails, "v", value)
	}
}
