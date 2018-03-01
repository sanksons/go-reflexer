package main

import (
	"fmt"

	"github.com/sanksons/go-reflexer"
)

//The example illustrates how reflexer can be used
// to set values for a struct, using reflection.
func example2() {

	//Lets define an example struct.
	type User struct {
		Name  string
		Title string
		Age   int
	}
	//dummy data to put in struct
	dummydata := map[string]interface{}{
		"name":  "Rachel",
		"title": "Miss",
		"age":   30,
	}

	//lets define our variable, which needs to be populated with data.
	user := User{}

	//One important point to note here is that we can't directly pass the struct, As reflection does not allow to Set data
	//directly in struct.
	//Instead, we will need to pass a pointer to the struct.

	reflectObj := reflexer.ReflectObj{}
	reflectObj.Initiate(&user)
	if !reflectObj.CheckIfPtr() { //since we expect a pointer here, check for it.
		panic("Expects a pointer.")
	}
	if !reflectObj.HasChild() {
		panic("No more levels to drill")
	}
	childStruct := reflectObj.GetChild()
	if !childStruct.CheckIfStruct() {
		panic("Expects a struct")
	}

	//extract field info of struct.
	structInfo, _ := reflexer.GetInfoAboutFieldsofStruct(*childStruct)
	fmt.Printf("%v\n", structInfo)

	//Put dummy data in struct.
	for k, val := range dummydata {
		index, ok := structInfo[k]
		if !ok {
			fmt.Println("Field not found, skipping")
			continue
		}
		switch v := val.(type) {
		case int64:
		case int:
			v64 := int64(v)
			childStruct.V.FieldByIndex([]int{index}).SetInt(v64)
		case string:
			childStruct.V.FieldByIndex([]int{index}).SetString(v)
		default:
			fmt.Printf("I don't know about type %T!\n", v)
		}

	}

	fmt.Printf("%v\n", user)

}

func main() {
	example2()
}
