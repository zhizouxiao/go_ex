package main

import (
	"fmt"
	"reflect"
)

func test() (int, error) {
	fmt.Println("ttttttt")
	return 1, nil
}

func main() {
	//var t int = 2
	t := test
	var p = fmt.Printf
	//p("%+v \n", t.Type())
	p("%+v \n", reflect.ValueOf(t).Type())
	p("%+v \n", reflect.ValueOf(t).Kind())
	p("%+v \n", reflect.Indirect(reflect.ValueOf(t)))
	p("%+v \n", reflect.TypeOf(t))

	type MyFloat float64
	var x MyFloat = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))
	fmt.Println("value in:", reflect.ValueOf(x).Float())
	fmt.Println("value type:", reflect.ValueOf(x).Type())
	fmt.Println("type kin:", reflect.ValueOf(x).Kind())
	fmt.Println("value kind:", reflect.ValueOf(x).Kind())
	var m reflect.Type
	a := 1
	m = reflect.ValueOf(a).Type()
	fmt.Println(m)
	switch m.Kind() {
	case reflect.Int32, reflect.Int64:
		fmt.Println("=====")
	}

}
