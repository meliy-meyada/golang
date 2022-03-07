package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	meya := person{
		firstName: "Meyada",
		lastName:  "Saisan",
		contactInfo: contactInfo{
			email:   "meyada.info@gmail.com",
			zipCode: 94000,
		},
	}
	meyaPointer := &meya
	meyaPointer.updateName("mayrisa")
	meya.print()


}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
 
func (p person) print() {

	fmt.Printf("%+v", p)

}
