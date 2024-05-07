package main

import "fmt"

type person struct {
	name string
	age int
}

func (p *person) updateAge(newAge int){
	(*p).age = newAge
}

func (p person) showInfo() {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("===================")
}

func main() {
	pers1 := person{
		name: "Jack",
		age: 10,
	}

	pers1.showInfo()
	pers1.updateAge(55)
	pers1.showInfo()
}

//Output
//Name: Jack
//Age: 10
//===================
//Name: Jack
//Age: 55
//===================