package main

import "fmt"

type Hasname interface {
	getname() string
}

func test(hasname Hasname) {
	fmt.Println("helo", hasname.getname())
}

type Person struct {
	name string
	umur int
}

func (person Person) getname() string {
	return person.name
}

type Hewan struct {
	name string
	umur int
}

func (hewan Hewan) getname() string {
	return hewan.name
}

func main() {
	var jehan Person
	jehan.name = "GG"

	test(jehan)

	contoh := Hewan{
		name: "kucing",
	}
	test(contoh)
}
