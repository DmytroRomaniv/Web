package Controls

import (
	"fmt"
	"net/http"
)

func AddProducts(writer http.ResponseWriter, request *http.Request){
	//TODO: Save to database

	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm error: %v", err)
		return
	}


}