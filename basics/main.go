package main

import (
	"fmt"
	"github.com/pmanickam81/go-tutorial/basics/util"
)

func main() {
	//variables
	var name = "Manickam"
	var age int32 = 36
	var isCool = false
	fmt.Printf("The name is %v\n", name)
	fmt.Println("The name is", name, "and the age is", age)
	fmt.Printf("The type is %T\n", age)
	fmt.Println(name, age, isCool)
	fmt.Println(util.Reverse("Hello"))
	fmt.Println(greeting("World"))
	fmt.Println(getSum(3, 2))

	//Arrays
	var textArr [3]string
	textArr[0] = "A"
	textArr[1] = "B"
	textArr[2] = "C"
	fmt.Println(textArr)
	fmt.Println(textArr[2])
	fmt.Println(textArr[1:])
	fmt.Println(textArr[:3])

	//Slices
	textSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(textSlice)
	textSlice = append(textSlice, 6)
	fmt.Println(textSlice)
	fmt.Println(len(textSlice))

	//Conditionals
	x := 14
	y := 6
	if x < y {
		fmt.Printf("%d is less than %d \n", x, y)
	} else {
		fmt.Printf("%d is greater than %d \n", x, y)
	}

	//Loops
	for i := 1; i <= 20; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	//Map
	emails := make(map[string]string)
	emails["A"] = "a@a.com"
	emails["B"] = "b@b.com"
	emails["C"] = "c@c.com"
	fmt.Println(emails)
	fmt.Println(emails["A"])
	fmt.Println(len(emails))
	delete(emails, "C")
	fmt.Println(emails)
	//Another way of Creating Map
	email := map[string]string {"A" :"a@a.com", "B" :"b@b.com", "C" :"c@c.com"}
	fmt.Println(email)

	//Range
	ids := []int{1,2,3,4,5}
	for i, id := range ids{
		fmt.Printf("%d -id: %d\n", i, id)
	}
	//Not using index
	for _, id := range ids{
		fmt.Printf("id: %d\n", id)
	}
	sum := 0
	for _, id := range ids{
		sum += id
	}
	println(sum)

	//Range for Map
	for k,v := range emails{
		fmt.Printf("%s: %s\n", k, v)
	}
	for k := range emails {
		fmt.Println("Key: " + k)
	}

}

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int32, num2 int32) int32 {
	return num1 + num2
}
