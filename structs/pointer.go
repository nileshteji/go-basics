package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

// reciever function can only be declared to type
type slice []string

func main() {
	nilesh := person{"Nilesh", "Teji"}
	// this is so cool
	sliceOne := slice{"HI", "My", "Name", "is", "Nilesh"}
	sliceOne.update()
	oneString := "Nilesh"
	update(oneString)
	fmt.Println(oneString)
	fmt.Println(sliceOne[0])
	nilesh.print()
	fmt.Println(nilesh.firstName)

}

func update(s string) {
	s = "Teji"
}

// slice and array are different
func (a slice) update() {
	a[0] = "Bye"
}

func (pointer *person) print() {
	(*pointer).firstName = "Balu"
}
