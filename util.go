package main

import "fmt"
import "strings"

// returns a map of field names and field indexes
func GetInfoAboutFieldsofStruct(structObj ReflectObj) (map[string]int, error) {
   //check if its a struct
   if structObj.Kind != reflect.Struct {
       return nil, fmt.Errorf("Expected Struct type, Got %v instead", structObj.Kind)
   }
   fieldCount := structObj.T.NumField()
   fieldInfo := make(map[string]int)  
   for i := 0; i < fieldCount; i++ {
	field := structObj.T.Field(i)
	  fmt.Printf("name: %v, Tag: %v \n", field.Name, field.Tag)
	  fieldInfo[strings.ToLower(field.Name)] = i
   }
   return fieldInfo, nil

}

