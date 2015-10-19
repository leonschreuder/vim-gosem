package main

import "fmt"

var someFieldToPrint = "value of someFieldToPrint"
var someOtherField = "value of someOtherField"

func someFunc() {
	stringToPrint := "printed"
	fmt.Println(stringToPrint)
	fmt.Println(someFieldToPrint)
	fmt.Println(someOtherField)
}
