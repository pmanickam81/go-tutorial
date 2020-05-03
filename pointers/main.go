package main

import "fmt"

func main(){
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	fmt.Println("To read the value from b", *b)
	fmt.Println("Another way to read value is", *&a)

	//Change the value
	*b = 10 //Pointers use reference
	fmt.Println("b is:", a)
}
