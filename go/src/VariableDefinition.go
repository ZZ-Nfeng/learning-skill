package main

import (
	"fmt"
)

func main() {
	//其中也可以多重声明
	var (
		a        = 1
		b        = "张三"
		c string = "王五"
	)
	// 1：指定变量类型，声明后不赋值，默认为0
	var name string
	var age int

	// 2：指定变量类型，声明赋值
	var address string = "陕西西安"

	// 3：定义变量但不指定变量类型，根据值自行判断变量类型
	var phone = 18978902789
	/*
		Println：打印字符串、变量；

		Printf：打印需要格式化的字符串，可以输出字符串类型的变量；不可以输出整型变量和整型；

		简单理解，当需要格式化输出信息时，一般选择Printf，其余使用Println。
	*/

	//4：用定义赋值符号 :=  ,只能定义在方法当中，go语言中定义了变量就一定要使用，不然会报错
	value := 1314
	fmt.Println(value)

	//fmt.Println(name)
	fmt.Println(age)

	fmt.Println(address)

	fmt.Println(phone)

	fmt.Printf("%v", name)

	fmt.Printf("%v,%v,%v", a, b, c)
}
