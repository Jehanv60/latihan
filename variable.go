package main

import "fmt"

func main() {
	var nama string = "jehan"

	fmt.Println(nama)

	nama = "dwi"
	fmt.Println(nama)

	test := "variabel"
	fmt.Println(test)

	var umur int
	umur = 30
	fmt.Println(umur)

	pos := 17115
	fmt.Println(pos)

	var (
		kampus  = "gundar"
		kampus2 = "J1"
	)
	fmt.Println(kampus, kampus2)
	var (
		kampus3 = "UI"
		kampus4 = 334
	)
	fmt.Println(kampus3, kampus4)
}
