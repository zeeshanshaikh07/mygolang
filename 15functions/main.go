package main

import "fmt"

func main() {
	fmt.Println("welcome to functions")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes, proMessage := proAdder(2,4,5,6,7)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("Pro message is: ", proMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values{
		total += val
	}

	return total, "Hi Pro result function"
}

// func greeterTwo()  {
// 	fmt.Println("Another method")
// }

func greeter() {
	fmt.Println("Hello from golang")
}