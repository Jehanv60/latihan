package main

import "fmt"

func main() {
	orang := map[string]string{
		"nama":    "jehan",
		"jabatan": "gg",
	}
	fmt.Println(orang)

	orang["iseng"] = "ya"
	fmt.Println(orang)
	fmt.Println(orang["nama"])
	fmt.Println(orang["jabatan"])

	orang2 := make(map[string]string)
	orang2["nama"] = "jehan"
	orang2["jabatan"] = "gg"
	orang2["test"] = "ilok"
	fmt.Println(orang2)
	delete(orang2, "test")
	fmt.Println(orang2)

}
