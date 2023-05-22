package main

import (
	"fmt"
)

func main() {

	a, b := 10, 20
	fmt.Println(a, b)
	swap2(a, b)
	fmt.Println(a, b)

	swap3(&a, &b)
	fmt.Println(a, b)

}

// go函数可以有多个返回值
func swap1(a, b string) (string, string) {
	return b, a
}

// 值传递
func swap2(x, y int) {
	temp := 0
	temp = x
	x = y
	y = temp
}

//引用传递

func swap3(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
