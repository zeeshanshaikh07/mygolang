package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	zeeshan := User{"Zeeshan", "Zeeshan@gmail.com", true, 22}
	fmt.Println(zeeshan)
	fmt.Printf("Zeeshan details are: %+v\n", zeeshan)
	fmt.Printf("Name is %v and email is %v\n", zeeshan.Name, zeeshan.Email)
	zeeshan.GetStatus()
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}

func (u User) GetStatus()  {
	fmt.Println("Is user active: ", u.Status)
}