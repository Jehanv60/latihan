package main

import (
	"fmt"
)

type identitas struct {
	kota, propinsi, negara string
}

func changecountry(alamat *identitas) {
	alamat.negara = "Amerika"
}

func main() {
	asal1 := identitas{
		kota:     "bekasi",
		propinsi: "jawa barat",
		negara:   "indonesia",
	}
	asal2 := &asal1
	asal3 := &asal1
	asal4 := new(identitas)
	asal4.kota = "denpasar"
	asal4.propinsi = "bali"
	asal2.kota = "jakarta"
	*asal2 = identitas{
		kota:     "malang",
		propinsi: "jawa timur",
		negara:   "indonesia",
	}

	fmt.Println(asal1)
	fmt.Println(asal2)
	fmt.Println(asal3)
	fmt.Println(asal4)

	alamat := identitas{
		kota:     "malang",
		propinsi: "jawa timur",
		negara:   "Indonesia",
	}
	changecountry(&alamat)
	fmt.Println(alamat)

}
