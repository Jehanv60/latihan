package main

import "fmt"

func main() {
	contohpanic(true)
}

func contohdefer() {
	message := recover()
	if message != nil {
		fmt.Println("pesan :=", message)
	}
	fmt.Println("test defer berhasil")
}

func contohpanic(error bool) {
	defer contohdefer()
	// defer func() {							cara kedua pada defer
	// 	message := recover()
	// 	if message != nil {
	// 		fmt.Println("pesan :=", message)
	// 	}
	// }()

	if error {
		panic("aplikasi error")
	}

	fmt.Println("bisa nih")
}
