package main

import (
	"./Configuration"
	"./Controls"
	"fmt"
)

func main() {
	var server Configuration.Server

	server.AddPage("/", Controls.Authorize)

	fmt.Println("The server has started.")

	server.StartServer(nil)

	//http.HandleFunc("/", Controls.CreateAccount)
	//http.ListenAndServe("localhost:8080", nil)

}
