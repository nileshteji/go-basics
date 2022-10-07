package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args

	myFile, er := os.Open(args[1])
	if er != nil {
		fmt.Println(er)
		os.Exit(1)
	}
	io.Copy(os.Stdout, myFile)

}
