package Controlers

import (
	"fmt"
	"net/http"
)

type Message string
func (msg Message) ServeHTTP(response http.ResponseWriter, request *http.Request){
	fmt.Fprint(response, msg)
}

func AddStartPage(){
	AddPage("/", "View/Index.html")
}

func AddPage(pattern string, root string){
	http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request){
		http.ServeFile(writer, request, root)
	})
}