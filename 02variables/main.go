package main

import "fmt"

const LoginToken string = "hjksadhjk"

func main() {
	var username string = "Zeeshan"
	fmt.Println(username)
	fmt.Printf("variable is of type : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type : %T \n", isLoggedIn)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("variable is of type : %T \n", smallval)

	var smallFloat float64 = 255.45454545454
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type : %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type : %T \n", anotherVariable)

	var anotherVariable1 string
	fmt.Println(anotherVariable1)
	fmt.Printf("variable is of type : %T \n", anotherVariable1)

	// implicit type
	var website = "zeeshan.com"
	fmt.Println(website)

	// no var style
	numberOfUser := 30000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type : %T \n", LoginToken)

}
