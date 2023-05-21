package main

import "fmt"

func main() {

	const AddressName string = "陕西西安"
	const value = "张三"

	const (
		a, b = iota + 1, iota + 2
		c, d
	)

	fmt.Println(AddressName)
	fmt.Println(a + b + c)
	fmt.Printf("%v+%v", a+b, c)

}
