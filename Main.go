package main

import (
	"./Controlers"
	"fmt"
	"net/http"
)

func main() {
	Controlers.AddStartPage()

	fmt.Println("The server has started.")

	http.ListenAndServe("localhost:8080", nil)

}
