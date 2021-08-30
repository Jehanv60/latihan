package main

import "fmt"

func main() {
	multiname()
	firstn, _, lastn := multiname()
	fmt.Println(firstn)
	// fmt.Println(midn)
	fmt.Println(lastn)

}

func multiname() (string, string, string) {
	return "jehan", "gg", "wp"
}
