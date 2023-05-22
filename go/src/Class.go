package main

import "fmt"

type Hero struct {
	name string
	age  int
}

func (this *Hero) GetName() string {
	return this.name
}

func (this *Hero) SetName(name string) {
	this.name = name
}

func (receiver *Hero) Show() {

}
func main() {
	hero := Hero{"zhang", 1}
	fmt.Println(hero)
	hero.SetName("张三")
	name := hero.GetName()
	fmt.Println(name)
}
