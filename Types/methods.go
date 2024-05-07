package main

import "fmt"

type mile float32
type kilometr float32

func (m mile) convertToKM() kilometr {
	return kilometr(m * 1.609);
}

func (k kilometr) convertToMile() mile {
	return mile(k / 1.609);
}

func main() {
	var dist1 mile = 10
	var dist2 kilometr = 200

	dist2 = dist1.convertToKM()
	fmt.Println(dist1, dist2)

	dist1 = dist2.convertToMile()
	fmt.Println(dist1, dist2)
}

//Output
//10 16.09
//10 16.09