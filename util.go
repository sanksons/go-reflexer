package reflexer

import "fmt"
import "strings"
import "reflect"

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
		fieldInfo[strings.ToLower(field.Name)] = i
	}
	return fieldInfo, nil

}
