package main

import "fmt"

func main() {
	segitiga()
	segitigasiku()
	segitigasiku2()
	genapganjil()
	penjumlahanganjil()
	penjumlahanganjil2()
	fizzbuzz()
	bilanganprima()
	fibonacci(10)
	fibonacci(16)
	fibonacci(100)
}
func segitiga() {
	for i := 0; i <= 5; i++ {
		for j := 5; j >= i; j-- {
			fmt.Print(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Print("* ")
		}
		fmt.Println()

	}

}
func segitigasiku() {
	for i := 0; i <= 5; i++ {
		// fmt.Print(i)
		for j := 5; j >= i; j-- { //for j := i; j < 5; j++(cara kedua)
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func segitigasiku2() {
	for i := 0; i <= 5; i++ {
		// fmt.Print(i)
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func genapganjil() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			fmt.Println("angka ganjil", i)
		} else if i%2 == 0 {
			fmt.Println("angka genap", i)
		}

	}

}

func bilanganprima() {
	for i := 1; i <= 15; i++ {
		a := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				a++
			}
		}
		if a == 2 {
			fmt.Println(fmt.Sprintf("deret angka prima = %d", i))
		}
	}
}

func penjumlahanganjil() {
	temp := 0
	for i := 3; i <= 20; i++ {
		if i%2 == 1 && temp == 0 {
			temp = i + i
			fmt.Println(fmt.Sprintf("angka awal adalah %d + %d = %d", i, i, temp))
		} else if i%2 == 1 {
			temp = temp + i
			fmt.Println(fmt.Sprintf("penjumlahan %d + %d = %d", temp-i, i, temp))
		}

	}
}

func penjumlahanganjil2() {
	temp := 3
	for i := 3; i <= 20; i++ {
		if i%2 == 1 {
			temp = temp + i
			fmt.Println(fmt.Sprintf("penjumlahan %d + %d = %d", temp-i, i, temp))
		}

	}
}

func fizzbuzz() {
	for i := 0; i <= 15; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(fmt.Sprintf("Angka %d adalah fizz & buzz", i))
		} else if i%3 == 0 {
			fmt.Println(fmt.Sprintf("Angka %d adalah fizz", i))
		} else if i%5 == 0 {
			fmt.Println(fmt.Sprintf("Angka %d adalah buzz", i))
		}
		// else {
		// 	fmt.Println(fmt.Sprintf("Angka %d bukan fizz dan buzz", i))
		// }
	}
}

func fibonacci(num int) {
	a := 0
	b := 1
	c := b
	fmt.Printf("deret fibonacci: %d %d", a, b)
	for true {
		c = b
		b = a + b
		if b >= num {
			fmt.Println()
			break
		}
		a = c
		fmt.Printf(" %d", b)
	}
}
