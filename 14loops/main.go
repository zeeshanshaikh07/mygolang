package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")

	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	fmt.Println(days)

	// for i:= 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days {
	// 	fmt.Printf("index is and value is %v\n", day)
	// }

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto zee
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}

		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

	zee: 
		fmt.Println("Jumping at zee.com")

}