package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// darko := person{"Darko", "Pranjic"}
	// darko := person{firstName: "Darko", lastName: "Pranjic"}
	// fmt.Println(darko)
	var darko = person{
		firstName: "Darko",
		lastName:  "Pranjic",
		contact: contactInfo{
			email:   "darkopranjic.4@gmail.com",
			zipCode: 52210,
		},
	}

	darko.updateName("Darkić")

	darko.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(firstName string) {
	p.firstName = firstName
}
