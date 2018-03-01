package main

import (
	"fmt"
	"reflect"

	"github.com/sanksons/go-reflexer"
)

//The example illustrates how reflexer can be used
// to set values for a slice of type []User, using reflection.
func example1() {

	//Lets define an example struct.
	type User struct {
		Name  string
		Title string
		Age   int
	}
	//lets define our variable, which needs to be populated with data.
	user := []User{}

	//One important point to note here is that we can't directly pass the slice, As reflection does not allow to Set data
	//directly in slice.
	//Instead, we will need to pass a pointer to the slice.

	reflectObj := reflexer.ReflectObj{}
	reflectObj.Initiate(&user)
	if !reflectObj.CheckIfPtr() { //since we expect a pointer here, check for it.
		panic("Expects a pointer.")
	}
	//fmt.Println(reflectObj.String())
	childSlice, err := reflectObj.GetChild()
	if err != nil {
		panic(err.Error())
	}
	//fmt.Println(childSlice.String())
	if !childSlice.CheckIfSlice() { //since we expect a slice here check for it.
		panic("Expects a slice.")
	}

	childStruct, err := childSlice.GetChild()
	if err != nil {
		panic(err.Error())
	}
	if !childStruct.CheckIfStruct() { //since we expect a struct here check for it.
		panic("Expects a struct.")
	}
	//fmt.Println(childStruct.String())

	//Since we have reached struct lets extract structs field information.
	structInfo, err := reflexer.GetInfoAboutFieldsofStruct(*childStruct)
	fmt.Printf("%v\n", structInfo)

	//No, we have to create new structs of the type User and assign values to it.
	//Below code creates two such structs and assign values to it.
	newstructitem := reflect.New(childStruct.T)
	newstructitem.Elem().FieldByIndex([]int{0}).SetString("Raven")
	newstructitem.Elem().FieldByIndex([]int{1}).SetString("Miss")
	newstructitem.Elem().FieldByIndex([]int{2}).SetInt(23)

	newstructitem2 := reflect.New(childStruct.T)
	newstructitem2.Elem().FieldByIndex([]int{0}).SetString("Marcus")
	newstructitem2.Elem().FieldByIndex([]int{1}).SetString("Mr")
	newstructitem2.Elem().FieldByIndex([]int{2}).SetInt(52)

	//Now, Append both the structs to our slice
	newval := reflect.Append(childSlice.V, newstructitem.Elem(), newstructitem2.Elem())
	//IMPORTANT!!
	// we also need to call set on slice, to overwrite with new slice.
	childSlice.V.Set(newval)

	fmt.Printf("%v\n", user)
	//o/p: [{Raven Miss 23} {Marcus Mr 52}]

}

func main() {
	example1()
}
