package main

import "fmt"

func main() {
	result := getname("gg")
	fmt.Println(result)
	fmt.Println(getname("xx"))
}

func getname(name string) string {
	if name == "gg" {
		return "ggwp"

	} else {
		return "hello" + name
	}

}
