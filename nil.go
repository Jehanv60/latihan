package main

import "fmt"

func main() {
	coba := testmap("jehan")
	if coba == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(coba)
	}
}

func testmap(nama string) map[string]string {
	if nama == "" {
		return nil
	} else {
		return map[string]string{
			"nama": nama,
		}
	}
}
