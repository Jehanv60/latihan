package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "jehan"
	names[1] = "dwi"
	names[2] = "fahrian"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		91,
		95,
	}
	fmt.Println(values)
	fmt.Println(len(values))
}
