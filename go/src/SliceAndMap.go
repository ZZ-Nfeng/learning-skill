package main

import "fmt"

func main() {

	var nums []int
	nums[0] = 1

	strings := make([]string, 10, 100)
	strings[0] = "1"

	var s = []int{1, 2, 3}
	fmt.Println(s)

	var maps map[string]string
	maps["张"] = "sads1"

	m := make(map[string]string)
	m["1"] = "sadada"

	var mmap = make(map[string]string)
	mmap["1"] = "1"

	mmm := map[int]int{
		1: 2,
		2: 9,
	}
	i := len(mmm)
	fmt.Println(i)

}
