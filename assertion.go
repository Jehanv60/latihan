package main

import (
	"fmt"
)

func main() {
	gas := random()
	gasubah := gas.(string)
	fmt.Println(gasubah)

	switch value := gas.(type) {
	case string:
		fmt.Println("ini string", value)
	case int:
		fmt.Println("ini int", value)
	default:
		fmt.Println("gtw", value)
	}
}

func random() interface{} {
	return "Jehan"
}
