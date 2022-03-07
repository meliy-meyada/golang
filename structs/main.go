package main

import "fmt"



type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string 
	lastName string
	contact contactInfo
}


func main() {
	meya := person{
		firstName: "Meyada",
		lastName: "Saisan",
		contact: contactInfo{
			email: "meyada.info@gmail.com",
			zipCode: 94000,
		},
	}
	
	fmt.Printf("%+v", meya)

}