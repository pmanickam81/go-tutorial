package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	//fName string
	//lName string
	//city string
	//gender string
	//age int
	// Another way of declaring struct
	fName, lName, city, gender string
	age                        int
}

//Value Receiver function
func (p Person) greet() string {
	return "Hello, my name is " + p.fName + " " + p.lName + " and I am " + strconv.Itoa(p.age)
}

// Pointer Receiver function
func (p *Person) increaseAge() {
	p.age++
}

// Pointer Receiver function
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lName = spouseLastName
	}
}

func main() {
	person1 := Person{
		fName:  "John",
		lName:  "Smith",
		city:   "NewYork",
		gender: "M",
		age:    13,
	}
	fmt.Println(person1)
	fmt.Println(person1.fName)

	person2 := Person{"Anna", "Johnson", "New York", "F", 30}
	fmt.Println(person2.greet())

	person1.increaseAge()
	fmt.Println(person1.greet())

	person2.getMarried("Thompson")
	fmt.Println(person2.greet())
}
