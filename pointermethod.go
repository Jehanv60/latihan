package main

import "fmt"

func main() {
	variable := Man{"Jehan", "Narogong"}
	variable.siapa()
	fmt.Println(variable)
}

type Man struct {
	nama, alamat string
}

func (identitas *Man) siapa() {
	identitas.nama = "Tuan " + identitas.nama + " tinggal di "
}
