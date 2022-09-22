package server

import (
	"gp/getdata"
	"log"
	"net/http"
	"text/template"
)

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	data := getdata.GetArtistById(4)
	t.Execute(w, data)
}
