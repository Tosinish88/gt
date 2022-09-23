package server

import (
	"fmt"
	"gp/binddata"
	"gp/getdata"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	var data []binddata.FullData

	// var i int
	switch r.URL.Path {
	case "/":
		data = getdata.GetData()
		for i, _ := range data {
			t.Execute(w, data[i])
		}
	case "/artists/Id":
		id, err := strconv.Atoi(r.URL.Path[9:])
		fmt.Println(id)
		if err != nil {
			log.Fatalln(err)
		}
		data1 := getdata.GetArtistById(id)
		fmt.Fprintln(w, data1)

	}
}

// func ServerHandler(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("templates/index_start.html")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	t.Execute(w, nil)

// 	var data []binddata.FullData
// 	data = getdata.GetData()
// 	for i, _ := range data {
// 		artT, err := template.ParseFiles("templates/artistThumb.html")
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		artT.Execute(w, data[i])
// 	}

// 	t, err = template.ParseFiles("templates/index_end.html")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	t.Execute(w, nil)
// }
