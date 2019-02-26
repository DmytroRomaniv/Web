package main

import (
	"./Configuration"
	"./Controls"
	"fmt"
)

func main() {
	var server Configuration.Server

	server.AddLayout("/static/", "./View/css")
	server.AddLayout("/images/", "./View/images")
	server.AddPage("/", Controls.Item)

	fmt.Println("The server has started.")

	server.StartServer(nil)

	//http.HandleFunc("/", Controls.CreateAccount)
	//http.ListenAndServe("localhost:8080", nil)
}
