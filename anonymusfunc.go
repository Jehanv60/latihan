package main

import "fmt"

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}
	register("admin", blacklist)
	register("jehan", blacklist)
	//cara kedua
	register("admin", func(name string) bool {
		return name == "admin"
	})
	register("jehan", func(name string) bool {
		return name == "root"
	})
}

type coret func(string) bool

func register(name string, blacklist coret) {
	if blacklist(name) {
		fmt.Println("ditolak", name)
	} else {
		fmt.Println("masuk", name)
	}
}
