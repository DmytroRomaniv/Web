package Controls

import (
	"../Models"
	"html/template"
	. "net/http"
)

func AddItem(writer ResponseWriter, request *Request) {
	//TODO: Check for existance


}

func Item(writer ResponseWriter, request *Request) {
	viewData := Models.ViewData{
		Items: []Models.Item{
			Models.Item{
				Id:             0,
				Index:          0,
				Name:           "Name1",
				FullName: 		"Full Name 1",
				Price:          1.1,
				NumbersInStock: 1,
				Description:    "Description1",
				ImageRoot:		"images/8Pack.jpg",
			},
			Models.Item{
				Id:             1,
				Index:          1 ,
				Name:           "Name2",
				FullName: 		"Full Name 2",
				Price:          2.1,
				NumbersInStock: 1,
				Description:    "Description2",
			},
			Models.Item{
				Id:             2,
				Index:          0,
				Name:           "Name3",
				FullName: 		"Full Name 3",
				Price:          3.1,
				NumbersInStock: 1,
				Description:    "Description3",
			},
		},
	}

	tmpl, _ := template.ParseFiles("View/Index.html")
	tmpl.Execute(writer, viewData)
}