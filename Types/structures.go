package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	var tomas person = person{"Tomas", 42}
	var tomas2 = person{age: 15, name:"Tomas2"}
	tomas3 := person{name: "Tomas3"} //age default = 0

	fmt.Println(tomas.name, tomas.age)
	fmt.Println(tomas2.name, tomas2.age)
	fmt.Println(tomas3.name, tomas3.age)

	fmt.Println("===========================")

	tomas3.age = 2
	fmt.Println(tomas3.name, tomas3.age)

	fmt.Println("===========================\nPointers:")

	var pers *person = &tomas3;
	fmt.Println(pers.name, pers.age)
	pers = &tomas
	fmt.Println(pers.name, pers.age)

	var presonAgePointer *int = &pers.age
	fmt.Println(*presonAgePointer)
	*presonAgePointer = 50
	fmt.Println(*presonAgePointer)

}

//Output
//Tomas 42
//Tomas2 15
//Tomas3 0
//===========================
//Tomas3 2
//===========================
//Pointers:
//Tomas3 2
//Tomas 42
//42
//50