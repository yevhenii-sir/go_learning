package main

import "fmt"

func sumArrFirst(arr [5]int) int {
	var sum int = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i];
	}

	return sum
}

func sumArrSecond(arr [5]float32) (sum float32) {
	for _, value := range arr {
		sum += value
	}

	return
}

func arrIntFindMax(arr [5]int) (maxNumberValue int, maxNumberIndex int) {
	if (len(arr) > 0) {
		maxNumberValue = arr[0]
		maxNumberIndex = 0

		for i := 1; i < len(arr); i++ {
			if (maxNumberValue < arr[i]) {
				maxNumberValue = arr[i]
				maxNumberIndex = i
			}
		}
	}

	return
}

func main() {
	var intNumbersArray [5]int = [5]int{ 1, 2, 100, 4, 5 }
	float32NumbersArray := [5]float32{ 1.5, 3.7, 4.8, 6.1, 7.22}

	fmt.Println(sumArrFirst(intNumbersArray))
	fmt.Println(sumArrSecond(float32NumbersArray))

	maxIntVal, maxIntIndex := arrIntFindMax(intNumbersArray)
	fmt.Println("maxIntVal:", maxIntVal, "maxIntIndex:", maxIntIndex)
}

//Output
//112
//23.32
//maxIntVal: 100 maxIntIndex: 2