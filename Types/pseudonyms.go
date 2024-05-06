package main

import "fmt"

type mile = int
type kilometr = int

func main() {
	var dist1 mile = 10
	var dist2 kilometr = 200

	dist1 = dist2 //Without error. its pseudonyms
	fmt.Println(dist1, dist2)
}

//Output
//200 200