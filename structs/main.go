package main

import "fmt"

type card struct {
	suit  string
	value string
}

/* So what actually we are doing here is



we get a address of a struct from &(amp percent) dont pronounce it and please.....
example:  cardPointer := &card
this cardPointer is the reference for the  card struct



we can get the value of the addess from using * on the address reference or pointer...
we can  get value of that pointer using this
(*cardPointer).suit this is actually the value of that reference




*/

func main() {
	nileshCard := card{"Spades", "value"}
	nileshPointer := &nileshCard
	nileshPointer.print()
	fmt.Println((*nileshPointer).suit)

}

// & operator working : give me the memory address of the value this variable is pointing at
// whenever we see  * before a type then it is not a operator (Desciption)

func (pointer *card) print() {
	(*pointer).suit = "Nilesh"
}
