/*
2. Declare two variables, `firstName` and `lastName` assign them with wanted values.
Declare constant named `fullname`
Combine constant and both strings into a full name by concatenating strings with a space in between and print them out.
*/
package main

import (
	"fmt"

)

func main()  {
	var firstName string = "John"
	var lastName string = "Doe"

	//firstName + " " + lastName (value of type string) is not constant
	//Promijenio sam u varijablu jer ne može biti konstanta iz razloga jer konstanta mora biti fiksna veličina 
	var fullName string= firstName +" "+ lastName
	

	fmt.Println(fullName)
	
}