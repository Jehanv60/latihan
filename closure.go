package main

import "fmt"

func main() {
	angka := 0
	nama := "yoman"

	increment := func() {
		nama := "GGWP"
		fmt.Println("test closure")
		angka++
		fmt.Println(nama)
	}

	increment()
	increment()

	fmt.Println(angka)
	fmt.Println(nama)

}
