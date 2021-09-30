package main

import "fmt"

func main() {

	// var color map[string]string
	colors := make(map[string]string)

	colors["nilesh"] = "teji"
	colors["nilesh"] = "taheem"
	colors["mrinank"] = "kumar"
	colors["vineet"] = "kumar"
	colors["savita"] = "taheem"
	colors["Ram"] = "Ram"

	delete(colors, "nilesh")

	for index, values := range colors {
		fmt.Println(index + " " + values)
	}
	// values := map[string]string{
	// 	"red": "#FFF",
	// }

}
