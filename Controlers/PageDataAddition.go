package Controlers

import (
	"html/template"
	"net/http"
)

func AddProducts(writer http.ResponseWriter, request *http.Request){

	tmpl, _ := template.ParseFiles("View/Index.html")

	tmpl.Execute(writer, data)
}