package main

import "fmt"

type a struct {
	a, b int
}

func (self a) method1() {
	self.a = 1
	self.b = 1
	fmt.Println(self)
}

func (self *a) method2() {
	self.a = 1
	self.b = 1
	fmt.Println(*self)
}
func main() {
	instance := a{0, 0}
	instance.method1()
	fmt.Println(instance)
	instance.method2()

	fmt.Println(instance)

}
