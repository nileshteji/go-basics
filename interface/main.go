package main

import "fmt"

// this is decalred interface and whoever have this method which is included in interface is also a honory member of
//of the intterfaces
type bot interface {
	// this need to be always a reciver function
	getGreeting() string
}
type englishBot struct {
}
type spanishbot struct {
}

func main() {

	eb := englishBot{}
	sb := spanishbot{}
	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// intterfaces are mainly used when we want reuse one function again again which is  same for 2 or more types of struct
func (eng englishBot) getGreeting() string {
	return "Hi Go!"
}

func (spanish spanishbot) getGreeting() string {
	return "Hola"
}
