package main

import "fmt"

func main() {
	name := "12"

	switch name {
	case "jehan":
		fmt.Println("hello jehan")
	case "joko":
		fmt.Println("hello Joko")
	default:
		fmt.Println("GG")

	}
	length := len(name)

	switch length >= 5 {
	case true:
		fmt.Println("nama lebih")
	case false:
		fmt.Println("EA")

	}

	switch length := len(name); length >= 5 {
	case true:
		fmt.Println("nama lebih")
	case false:
		fmt.Println("EA")

	}

	switch {
	case length >= 10:
		fmt.Println("nama lebih")
	case length >= 5:
		fmt.Println("nama mayan")
	default:
		fmt.Println("PRO")

	}
}
