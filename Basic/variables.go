package main

import "fmt"

func main() {
	//variables
	var a int = 5
	b := 2
	var resultString string = "Result:"
	var (
		x = 54.0000014
		y = 66
	)
	var v1, v2 = "variable1", "variable2"

	//output method 1
	fmt.Println(resultString, "a + b =", a, "+", b, "=", a + b)

	//output method 2
	resultString = fmt.Sprintf("%s a + b = %d + %d = %d", resultString, a, b, a + b)
	fmt.Println(resultString)

	resultString = fmt.Sprintf("x: %.2f, y: %d, v1: %s, v2: %s", x, y, v1, v2)
	fmt.Println(resultString)
}

//Output
//Result: a + b = 5 + 2 = 7
//Result: a + b = 5 + 2 = 7
//x: 54.00, y: 66, v1: variable1, v2: variable2