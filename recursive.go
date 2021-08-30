package main

import "fmt"

func main() {
	hasil := forloop(5)
	fmt.Println(hasil)
	hasil2 := recursive(5)
	fmt.Println(hasil2)
}

func forloop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i //sama dengan result = result * 1
	}
	return result
}

func recursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * recursive(value-1)
	}
}
