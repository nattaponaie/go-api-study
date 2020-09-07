package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func firstAndLast(str string) (string, string) {
	return string(str[0]), string(str[1])
}

func rectangle(width int, height int) (w int, h int) {
	w = width
	h = height
	return w, h
}

// Variadic Functions
func sum(numbers ...int) int {
	fmt.Println(numbers)
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func init() {
	fmt.Println("Init main...")
}

func main() {
	res1 := plus(1, 2)
	fmt.Println(res1)

	res201, res202 := firstAndLast("My name is Foo")
	fmt.Println(res201, res202)

	res301, res302 := rectangle(10, 20)
	fmt.Println(res301, res302)

	res4 := sum(12, 20, 1, 2)
	fmt.Println(res4)
}
