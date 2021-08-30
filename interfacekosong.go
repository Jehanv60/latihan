package main

import "fmt"

func main() {
	kosong := test("adasdas")
	fmt.Println(kosong)
}

func test(i string) interface{} {
	if i == "jehan" {
		return i
	} else if i == "bagen" {
		return 4
	} else {
		return "nama"
	}

}
