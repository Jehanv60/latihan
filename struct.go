package main

import "fmt"

type Customer struct {
	Name, Alamat string
	Umur         int
	Married      bool
}

func (customer Customer) sayhello() { //struct method
	fmt.Println("nama saya : ", customer.Name, "dari", customer.Alamat)
}

func (test Customer) sayhi(name string) { //struct method
	fmt.Println("nama saya : ", name, "dari", test.Alamat)
}
func main() {

	var jehan Customer
	jehan.Name = "jehan"
	jehan.Alamat = "narogong"
	jehan.Umur = 24
	jehan.Married = false

	fmt.Println(jehan)
	jehan.sayhello()
	jehan.sayhi("Jehan")
	// gg := Customer{
	// 	Name:    "han",
	// 	Alamat:  "rawalumbu",
	// 	Umur:    24,
	// 	Married: false,
	// }
	// fmt.Println(gg)

	// wp := Customer{"budi", "doremi", 40, true}
	// fmt.Println(wp)
}
