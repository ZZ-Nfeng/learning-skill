package main

import "fmt"

func main() {
	var animal Animal
	animal = &Cat{"red"}
	animal.Sleep()
	catType := animal.GetType()
	fmt.Println(catType)
	animal = &Dog{"dog"}
	animal.Sleep()
}

type Animal interface { //接口的数据类型是指针
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

func (c Cat) Sleep() {
	//TODO implement me
	panic("implement me")
}

func (c Cat) GetColor() string {
	//TODO implement me
	panic("implement me")
}

func (c Cat) GetType() string {
	//TODO implement me
	panic("implement me")
}

//func (this Cat) Sleep() {
//	fmt.Println("cat is sleep")
//}
//
//func (this Cat) GetColor() string {
//	return this.color
//}
//func (this Cat) GetType() string {
//	return this.color
//}

type Dog struct {
	color string
}

func (this Dog) Sleep() {
	fmt.Println("Dog is sleep")
}

func (this Dog) GetColor() string {
	return this.color
}
func (this Dog) GetType() string {
	return this.color
}
