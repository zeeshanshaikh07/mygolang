package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	// fruitList[0] = "Apple"
	fruitList[3] = "Grape"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [3]string {"potato", "beans", "tomato"}

	fmt.Println("VegList is:" , vegList)
	fmt.Println("VegList is:" , len(vegList))
}