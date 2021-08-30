package main

import "fmt"

func main() {
	total := sumall(10, 20, 30)
	fmt.Println(total)

	slice := []int{10, 40, 50}
	total = sumall(slice...)
	fmt.Println(total)
}

func sumall(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}
