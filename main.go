package main

import (
	"fmt"
	"sort"
)

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
	// age := 35
	// name := "Shaun"
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
	// var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	// fmt.Println("the saved string is:", str)

	// arrays (length is fixed)
	// var ages [3]int = [3]int{20, 25, 30}
	// var ages = [3]int{20, 25, 30}
	// fmt.Println(ages, len(ages))
	// names := [4]string{"yoshi", "mario", "peach", "bowser"}
	// names[1] = "luigi"
	// fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	// var scores = []int{100, 50, 60}
	// scores[2] = 25
	// scores = append(scores, 85)
	// fmt.Println(scores, len(scores))

	// slice ranges
	// rangeOne := names[1:3] // mario peach
	// rangeTwo := names[2:] // peach bowser
	// rangeThree := names[:3] // yoshi mario peach
	// fmt.Println(rangeOne, rangeTwo, rangeThree)
	// fmt.Printf("the type of rangeOne is %T \n", rangeOne) // []string
	// rangeOne = append(rangeOne, "koopa")
	// fmt.Println(rangeOne) // mario peach koopa

	// standard library
	// https://golang.org/pkg/
	
	// strings package - doesn't alter original package
	// greeting := "Hello there friends!"
	// fmt.Println(strings.Contains(greeting, "Hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "ll"))
	// fmt.Println(strings.Split(greeting, " "))

	// sort package
	// ages := []int{45, 20, 35, 30, 75, 68, 58, 25}
	// sort.Ints(ages)
	// fmt.Println(ages)
	// index := sort.SearchInts(ages, 30)
	// fmt.Println(index)
	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "bowser"))
}