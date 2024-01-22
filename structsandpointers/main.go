package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string //No comma
	lastName    string
	contactInfo //declares contactInfo contactInfo
}

func main() {

	/*
		Struct declaration type
	*/

	//alex := person{"Alex", "Anderson"}
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	// var alex person //zero value struct with default values ""
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	/*
		Embed one struct into another
	*/
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 123456,
		}, //comma needed for the last value
	}

	//Since Go uses pass by value approach we need to use pointers
	//& - give access to memory address of jim
	// jimPointer := &jim
	// jimPointer.updateName("jimmy")
	// jim.print()
	//Above 3 can be written as below and the function still works
	jim.updateName("jimmy")
	jim.print()
}

/*
Receiver
*/
func (p person) print() {
	fmt.Printf("%+v", p)
}

// *person - this is a type description: It means we are working with a pointer to a person
func (pointerToPerson *person) updateName(newFirstName string) {
	// *pointer - give value in this memory address. This is an operator - It means we want to manipulate the value pointer is referencing
	//* turn pointer into address
	(*pointerToPerson).firstName = newFirstName
}
