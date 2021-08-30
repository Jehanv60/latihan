package main

import "fmt"

func main() {
	name := "asd"

	if name == "jehan" {
		fmt.Println("GG")
	} else if name == "asd" {
		fmt.Println("surya")
	} else {
		fmt.Println("salah cuy")
	}

	if length := len(name); length >= 5 {
		fmt.Println("true")
	} else {
		fmt.Println("eeaaa")
	}
	test()
}

func test() {
	a := 0
	b := 1
	c := b //C = 1
	fmt.Printf("Series is: %d %d", a, b)
	for true {
		c = b
		b = a + b
		if b >= 16 {
			fmt.Println()
			break
		}
		a = c
		fmt.Printf(" %d", b)
	}
}
