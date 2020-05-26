package main

import "fmt"

type array []int

func main() {

	numbers := array{1, 2, 3, 4, 5, 6, 7, 8, 9, 20}

	for _, index := range numbers {
		if index%2 == 0 {
			fmt.Println(index, " is even")
		} else if index%2 != 0 {
			fmt.Println(index, "is Odd")
		}
	}

}
