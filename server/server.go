package server

import (
	"gp/getdata"
	"log"
	"net/http"
	"text/template"
	"gp/binddata"
)

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	var data []binddata.ArtistData
	switch r.URL.Path {
	case "/":
		data = getdata.GetData()
	// case "/artists/" + "id":
	// 	id := 
		
	}
	

	t.Execute(w, data)
}
