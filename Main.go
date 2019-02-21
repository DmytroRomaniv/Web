package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("The server has started.")

	http.ListenAndServe("localhost:8080", nil)

}
