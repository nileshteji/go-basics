package main

import "fmt"

func main() {
	cards := []int{1, 2, 3}
	for i, card := range cards {
		fmt.Println(i, card)
	}

}
