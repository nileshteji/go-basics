package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type customWriter struct{

}

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println(err)
		os.Exit(23)

	}


}
