package main

import "fmt"
import "reflect"

type square struct{ r int }
type circle struct{ r int }

func (s square) area() int { return s.r * s.r }
func (c circle) area() int { return c.r * 3 }

func main() {
	s := square{1}
	c := circle{1}
	a := [2]interface{}{s, c}
	fmt.Println(s, c, a)

	sum := 0
	for _, t := range a {
		fmt.Println(t)
		//fmt.Println(t.TypeOf())
		fmt.Println("type:", reflect.TypeOf(t))
		fmt.Println("type:", reflect.TypeOf(t))
		fmt.Println("type kin:", reflect.ValueOf(t).Kind())
		fmt.Println("type:", reflect.TypeOf(t).Kind())

		switch v := t.(type) {
		case square:
			sum += v.area()
		case circle:
			sum += v.area()
		}
	}
	fmt.Println(sum)
}
