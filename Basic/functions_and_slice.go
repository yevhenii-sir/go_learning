package main

import "fmt"

func sumInt32(numbers ...int32) (sum int32) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func main() {
	fmt.Println(sumInt32(1, 2, 3, 4, 5))

	//slice input
	var sliceData []int32 = []int32{ 1, 2, 3, 4, 5}
	fmt.Println(sumInt32(sliceData...))

	sliceData = append(sliceData, 100)
	fmt.Println(sumInt32(sliceData...))

	newSliceData := sliceData[2:5] //3..5 elements
	fmt.Println(newSliceData)
	fmt.Println(sumInt32(newSliceData...))
}

//Output
//15
//15
//115
//[3 4 5]
//12