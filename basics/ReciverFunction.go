package ReciverFunction

import "fmt"

type deck []string

func main() {

	cards := deck{"Nilesh", "teji"}
	cards.print()

	// will give compile type error
	//
	// card := "nilesh"
	// card.print()

}

// this means any variable of type deck can access this method
func (d deck) print() {

	for i := 0; i < len(d); i++ {
		fmt.Println(d[i])
	}

}
