package main

import "fmt"

func main() {
	hasil := getname
	result := hasil("jehan")
	fmt.Println(result)
	fmt.Println(getname("GG"))
}

func getname(nama string) string {
	return "hello" + nama
}
