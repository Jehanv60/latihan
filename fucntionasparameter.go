package main

import "fmt"

func main() {
	namafilter("jehan", spamfilter)
	namafilter("WP", spamfilter)
	result := spamfilter //cara kedua
	namafilter("yoman", result)
}

type Filter func(string) string

func namafilter(nama string, filter Filter) {
	namafiltered := filter(nama)
	fmt.Println("coba", namafiltered)
}

func spamfilter(nama string) string {
	if nama == "jehan" {
		return "GG"
	} else {
		return nama
	}
}
