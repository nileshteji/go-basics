package main

import "fmt"

type array []int
type deck []string

func main() {
	d := deck{"Nilesh", "Teji"}
	d.update()
	fmt.Println(d)
}

func (d deck) update() {
	d[0] = "Update"
}
