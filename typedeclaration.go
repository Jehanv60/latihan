package main

import "fmt"

func main() {
	type huruf string
	type angka int

	var nama huruf = "jehan"
	var umur angka = 24

	fmt.Println(nama)
	fmt.Println(umur)
}
