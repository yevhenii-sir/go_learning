package main

import "fmt"

type contact struct {
	email string
	phone string
}

type person struct {
	name string
	age int
	contactInfo contact
}

func main() {
	var tomas = person {
		name: "Tomas",
		age: 19,
		contactInfo: contact {
			email: "tipa_tomas@gmail.com",
			phone: "+380123456789",
		},
	}

	fmt.Println(tomas)
	fmt.Println(tomas.contactInfo.email)
    fmt.Println(tomas.contactInfo.phone)
}

//Output 
//{Tomas 19 {tipa_tomas@gmail.com +380123456789}}
//tipa_tomas@gmail.com
//+380123456789