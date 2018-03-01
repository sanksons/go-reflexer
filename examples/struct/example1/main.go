package main

import "./github.com/sanksons/go-reflexer"

func teststruct() {
	
	
		user := &[]User{}
		
		reflectObj := ReflectObj{}
		reflectObj.Initiate(user)
		if !reflectObj.CheckIfPtr() {
			panic("Expects a pointer.")
		}
		fmt.Println(reflectObj.String())
		childSlice, err := reflectObj.GetChild()
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(childSlice.String())
		
		childStruct, err := childSlice.GetChild()
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(childStruct.String())
		
		structInfo, err := GetInfoAboutFieldsofStruct(*childStruct)
		fmt.Printf("%v\n", structInfo)
		
		
		
		newstructitem := reflect.New(childStruct.T)
		fmt.Printf("kind: %s, kind: %s\n",newstructitem.Kind(),newstructitem.Kind())
		newstructitem.Elem().FieldByIndex([]int{0}).SetString("sdsdds")
		newstructitem.Elem().FieldByIndex([]int{1}).SetString("fgfgffggf")
		
		newval := reflect.Append(childSlice.V,newstructitem.Elem())
		
		childSlice.V.Set(newval)
		
		
		
		fmt.Printf("%v\n", user)
		
		
	}
	