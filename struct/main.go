package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	jin := person{
		firstname: "Patchara",
		lastname:  "Chukiatkajohn",
		contact: contactInfo{
			email: "test@ttmail.com",
			zip:   1234,
		},
	}
	// var jin person
	// jin.firstname = "Patchara"
	// jin.lastname = "Chukiatkajohn"

	// fmt.Println("%+v", jin)

	// jinPointer := &jin
	// jinPointer.updateFirstname("Jin")
	jin.updateFirstname("Jin")
	jin.printPerson()
}

func (p person) printPerson() {
	fmt.Println("%+v", p)
}

func (p *person) updateFirstname(name string) {
	(*p).firstname = name
}
