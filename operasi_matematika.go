package main

import "fmt"

func main() {
	var hasil = 10 + 10
	fmt.Println(hasil)

	var a = 100
	var b = 1
	var c = a + b
	c++
	fmt.Println(c)

	var temp int
	var d = 5
	var e = 3

	temp = d //5
	d = e    //3
	e = temp //5

	var f = 5
	var g = 3
	f = f * g //15
	g = f / g //5
	f = f / g //3

	fmt.Println("nilai d =", d)
	fmt.Println("nilai e =", e)
	fmt.Println("nilai f =", f)
	fmt.Println("nilai g =", g)
}
