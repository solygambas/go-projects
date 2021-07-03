package main

import "fmt"

func main() {
	// strings
	// var nameOne string = "Mario"
	// var nameTwo = "Luigi"
	// var nameThree string
	// fmt.Println(nameOne, nameTwo, nameThree)
	// nameOne = "Peach"
	// nameThree = "Bowser"
	// fmt.Println(nameOne, nameTwo, nameThree)
	// nameFour := "Yoshi"
	// fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	// integers 
	// var ageOne int = 20
	// var ageTwo = 25
	// ageThree := 30
	// fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	// https://golang.org/ref/spec#Numeric_types
	// var numOne int8 = 127
	// var numTwo int8 = -128
	// var numThree uint8 = 255
	// var numFour uint16 = 256

	// floats
	// var scoreOne float32 = 25.98
	// var scoreTwo float64 = 25.55555555555555555555555555
	// scoreThree := 25.98

	// Print
	// fmt.Print("hello, ")
	// fmt.Print("world! \n")
	// fmt.Print("new line \n")

	// Println
	// fmt.Println("new line")
	age := 35
	name := "Shaun"
	// fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted strings)
	// %_ = format specifier
	// https://golang.org/pkg/fmt/
	// fmt.Printf("my age is %v and my name is %v \n", age, name)
	// fmt.Printf("my age is %q and my name is %q \n", age, name) // quotes (only with strings)
	// fmt.Printf("age is of type %T", age) // type
	// fmt.Printf("you scored %f points! \n", 225.55) // float 225.550000
	// fmt.Printf("you scored %0.1f points! \n", 225.55) // float 225.6

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)
}