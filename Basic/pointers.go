package main

import "fmt"

func mathPower(value *int, n int) {
	startValue := *value;
	for n > 1 {
		*value = *value * startValue;
		n--
	}
}

func main() {
	var x int = 5
	var pointerX *int = &x

	if pointerX != nil {
		fmt.Println("Variable address: ", pointerX)
		fmt.Println("Variable data: ", *pointerX)

		mathPower(pointerX, 3)
		fmt.Println("Variable next data: ", x)
	}

	var newPointer *int = new(int)
	fmt.Println("Value:", *newPointer)
	*newPointer = 37 //:)
	fmt.Println("Value:", *newPointer)
}

//Output
//Variable address:  0xc000096068
//Variable data:  5
//Variable next data:  125
//Value: 0
//Value: 37