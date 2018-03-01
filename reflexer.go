package main

import "reflect"
import "fmt"


type ReflectObj struct {
	V reflect.Value
	T reflect.Type
	Kind reflect.Kind
	child *ReflectObj
	parent *ReflectObj
}

func (this *ReflectObj) String() string {
	return fmt.Sprintf("ReflectObj { V:%v, T:%T, K:%s}",this.V,this.T,this.Kind)
}	

func (this *ReflectObj) HasChild() bool {
  if this.CheckIfPtr() || this.CheckIfSlice() {
     return true
  }
  return false
}

//Make sure to call Has child before calling this.
func (this *ReflectObj) GetChild() (*ReflectObj, error) {
   	
   if this.child != nil {
	   return this.child, nil
   }
   child := new(ReflectObj)
   child.T = this.T.Elem() 
   if this.Kind == reflect.Ptr || this.Kind == reflect.Interface {
	child.V = this.V.Elem()
	}
   child.SetParent(this)	
   this.child = child
   return this.child, nil
}

func (this *ReflectObj) GetParent() (*ReflectObj) {
   return this.parent	
}

func (this *ReflectObj) SetParent(parent *ReflectObj) {
	this.parent = parent
 }


func (this *ReflectObj) Initiate(i interface{}) {
   
   this.V = reflect.ValueOf(i)
   this.T = this.V.Type()
   this.Kind = this.T.Kind()
}


func (this *ReflectObj) CheckIfPtr() bool {
	if this.Kind != reflect.Ptr {
		return false
	}
	return true
}

func (this *ReflectObj) CheckIfSlice() bool {
	if this.Kind != reflect.Slice {
		return false
	}
	return true
}

