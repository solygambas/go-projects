package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var score = 99.5

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func getInitials(n string) (string, string) {
 s := strings.ToUpper(n)
names := strings.Split(s, " ")
var initials []string
for _, v := range names {
	initials = append(initials, v[:1])
}

if len(initials) > 1 {
	return initials[0], initials[1]
}

return initials[0], "_"
}

func updateName(x string) string {
	x = "wedge" // local copy
	return x // return is needed to update original value
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func updateNameP(x *string) {
	*x = "wedge"
}

// interfaces
type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}
type circle struct {
	radius float64
}

// square methods
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumf() float64 {
	return s.length * 4
}

// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f \n", s, s.circumf())
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("item added - ", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("tip added - ", tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("you saved the file - ", b.name)
	default:
		fmt.Println("that was not a valid option...")
		promptOptions(b)
	}
}

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
	// names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	// sort.Strings(names)
	// fmt.Println(names)
	// fmt.Println(sort.SearchStrings(names, "bowser"))

	// for loop
	// x := 0
	// for x < 5 {
	// 	fmt.Println("value fo x is:", x)
	// 	x++
	// }
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value fo i is:", i)
	// }
	// names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }
	// for index, value := range names {
	// 	fmt.Printf("the value at index %v is %v \n", index, value)
	// }
	// for _, value := range names {
	// 	fmt.Printf("the value is %v \n", value)
	// }

	// booleans
	// age := 45
	// fmt.Println(age <= 50)
	// fmt.Println(age >= 50)
	// fmt.Println(age == 45)
	// fmt.Println(age != 50)

	// if age < 30 {
	// 	fmt.Println("age is less than 30")
	// } else if age < 40 {
	// 	fmt.Println("age is less than 40")
	// } else {
	// 	fmt.Println("age is not less than 40")
	// }

	// names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	// for index, value := range names {
	// 	if index == 1 {
	// 		fmt.Println("continuing at pos", index)
	// 		continue // skip code after
	// 	}
	// 	if index == 2 {
	// 		fmt.Println("breaking at pos", index)
	// 		break // stop the loop
	// 	}
	// 	fmt.Printf("the value at pos %v is %v \n", index, value)
	// }

	// functions
	// sayGreeting("Mario")
	// sayBye("Luigi")

	// cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)

	// a1 := circleArea(10.5)
	// a2 := circleArea(15)
	// fmt.Println(a1, a2)
	// fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)

	// returning multiple values
	// firstName, secondName := getInitials("Tifa Lockhart")
	// fmt.Println(firstName, secondName)
	// firstName2, secondName2 := getInitials("cloud strife")
	// fmt.Println(firstName2, secondName2)
	// firstName3, secondName3 := getInitials("barret")
	// fmt.Println(firstName3, secondName3)

	// package scope
	// sayHello("Mario")
	// for _, v := range points {
	// 	fmt.Println(v)
	// }
	// showScore()

	// maps
	// menu := map[string]float64{
	// 	"soup": 4.99,
	// 	"pie": 7.99,
	// 	"salad": 6.99,
	// 	"toffee pudding": 3.55,
	// }
	// fmt.Println(menu)
	// fmt.Println(menu["pie"])

	// looping maps
	// for k, v := range menu {
	// 	fmt.Println(k, "-", v)
	// }

	// ints as key type
	// phonebook := map[int]string{
	// 	267584967: "mario",
	// 	984759373: "luigi",
	// 	845775485: "peach",
	// }
	// fmt.Println(phonebook)
	// fmt.Println(phonebook[267584967])
	// phonebook[267584967] = "bowser"
	// fmt.Println(phonebook)
	// phonebook[845775485] = "yoshi"
	// fmt.Println(phonebook)

	// pass by value
	// non-pointer values: strings, ints, bools, floats, arrays, structs 
	// name := "tifa"
	// name = updateName(name)
	// fmt.Println(name)
	// pointer wrapper values: slices, maps, functions
	// menu := map[string]float64{
	// 	"pie": 5.95,
	// 	"ice cream": 3.99,
	// }
	// updateMenu(menu)
	// fmt.Println(menu)

	// pointers
	// name := "tifa"
	// updateName(name)
	// fmt.Println("memory address of name is:", &name) // 0xc000026250
	// m := &name // stores the pointer to 0xc000026250
	// // fmt.Println("memory address:", m) // 0xc000026250
	// // fmt.Println("value at memory address:", *m) // tifa
	// fmt.Println(name) // tifa
	// updateNameP(m)
	// fmt.Println(name) // wedge

	// structs
	// mybill := newBill("mario's bill")
	
	// receiver functions
	// mybill.addItem("onion soup", 4.50)
	// mybill.addItem("veg pie", 8.95)
	// mybill.addItem("toffee pudding", 4.95)
	// mybill.addItem("coffee", 3.25)
	// mybill.updateTip(10)
	// fmt.Println(mybill.format())

	// interfaces
	// shapes := []shape{
	// 	square{length: 15.2},
	// 	circle{radius: 7.5},
	// 	circle{radius: 12.3},
	// 	square{length: 4.9},
	// }

	// for _, v := range shapes {
	// 	printShapeInfo(v)
	// 	fmt.Println("---")
	// }

	// user input
	mybill := createBill()
	promptOptions(mybill)
}