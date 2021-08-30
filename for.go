package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println("perulangan ke", i)

	}

	var array = [4]string{"GG", "ez", "wp", "thanks"}

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])

	}

	for i, value := range array {
		fmt.Println(i, value)
		// fmt.Println(value) jika ingin menggunakan tanpa i maka i diganti _

	}

	person := make(map[string]string)
	person["nama"] = "jehan"
	person["jabatan"] = "gg"

	for j, gogo := range person {
		fmt.Println(j, gogo)

	}

	for i := 0; i <= 5; i++ {
		for j := 5; j >= i; j-- {
			fmt.Print("* ")

		}
		fmt.Println()
	}
}
