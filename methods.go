package main

import "fmt"

type User struct {
	firstName, lastName string 
	age int8
}

func (args User) CreateUser() string {
	usr := User{
		firstName: args.firstName,
		lastName: args.lastName,
		age: args.age,
	}

	str := fmt.Sprintf("I'm %v %v", usr.firstName, usr.lastName)
	return str
}

func main() {
	user := User{
		firstName: "Morteza",
		lastName: "Heydari",
		age: 24,
	}
	
	fmt.Println(user.CreateUser())
}