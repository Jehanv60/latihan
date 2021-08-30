package main

import "fmt"

func main() {
	var bulan = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	fmt.Println(bulan)

	var slice = bulan[4:7]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// bulan[4] = "GG"
	// fmt.Println(bulan)

	slice[0] = "expert"
	fmt.Println(bulan)

	var slice1 = append(slice, "eko")
	fmt.Println(slice1)

	slice5 := bulan[3:10]
	fmt.Println(slice5)

	var slice6 = append(slice5, "eko")
	fmt.Println(slice6)
	fmt.Println(slice)
	fmt.Println(bulan)

	var newslice = make([]string, 2, 5)
	newslice[0] = "jehan"
	newslice[1] = "dadas"

	fmt.Println(newslice)

	copyslice := make([]string, len(newslice), cap(newslice))
	copy(copyslice, newslice)
	fmt.Println(copyslice)

	slice200 := []string{"jehan", "vendor"}
	for i := 0; i < len(slice200); i++ {
		fmt.Println(slice200[i])
	}
}
