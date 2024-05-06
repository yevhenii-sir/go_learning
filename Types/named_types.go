package main

import "fmt"

type BinaryOp func(int, int) int
type mile int
type kilometr int

func action(a int, b int, op BinaryOp) {
	result := op(a, b)
	fmt.Println(result)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func main() {
	action(10, 5, add)
	action(10, 5, sub)

	var dist1 mile = 10
	var dist2 kilometr = 200

	//dist1 = dist2 //ERROR! its different types
	fmt.Println(dist1, dist2)
}

//Output
//15
//5
//10 200