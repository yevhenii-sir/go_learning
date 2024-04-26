package main

import "fmt"

func main() {
	var testRate = map[string]uint8 {
		"Rostislav": 12,
		"Yevhenii": 9,
		"Alice": 2,
	}

	//check exists
	if val, exists := testRate["Yevhenii"]; exists{
		fmt.Println("Key Yevhenii exists! Rate:",val)
	}

	//view all
	for key, value := range testRate{
		fmt.Println(key, "==>", value)
	}

	fmt.Println("===============")
	delete(testRate, "Rostislav");
	for key, value := range testRate{
		fmt.Println(key, "==>", value)
	}
}

//Output
//Key Yevhenii exists! Rate: 9
//Rostislav ==> 12
//Yevhenii ==> 9
//Alice ==> 2
//===============
//Yevhenii ==> 9
//Alice ==> 2