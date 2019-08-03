package main

import (
	"fmt"
	"reflect"
)

var typeRegistry = make(map[string]reflect.Type)

func registerType(elem interface{}) {
	t := reflect.TypeOf(elem).Elem()
	typeRegistry[t.Name()] = t
}

func newStruct(name string) (interface{}, bool) {
	elem, ok := typeRegistry[name]
	if !ok {
		return nil, false
	}
	return reflect.New(elem).Elem().Interface(), true
}

func init() {
	registerType((*test)(nil))
	registerType((*test1)(nil))
}

type test struct {
	Name string
	Sex  int
}

type test1 struct {
	T test
}

func main() {
	structName := "test1"

	s, ok := newStruct(structName)
	if !ok {
		return
	}

	fmt.Println(s, reflect.TypeOf(s))

	/* 	t, ok := s.(test1)
	   	if !ok {
	   		return
	   	}
	   	// t.Name = "i am test"
	   	fmt.Println(t, reflect.TypeOf(t)) */
}
