package main

import (
	"errors"
	"fmt"
)

func main() {
	hasil, err := cobaerror(100, 10)
	if err == nil {
		fmt.Println("hasil", hasil)
	} else {
		fmt.Println("gak bisa klu 0", err.Error())
	}
	valid, err := errorstring("GG", "T")
	if err == nil {
		fmt.Println("Nama", valid)
	} else {
		fmt.Println("nama beda", err.Error())
	}

}

func cobaerror(nilai int, nilai1 int) (int, error) {
	if nilai1 == 0 {
		return 0, errors.New("tidak bsia")
	} else {
		result := nilai / nilai1
		return result, nil
	}
}

func errorstring(nama string, nama2 string) (string, error) {
	if nama2 == "GG" {
		return nama, errors.New("nama tidak sesuai")
	} else {
		return nama, nil
	}

}
