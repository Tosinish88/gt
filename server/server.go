package server

import (
	"gp/getdata"
	"log"
	"net/http"
	"text/template"
)

var ArtistsData []getdata.FullData

// func contains(s []string, str string) bool {
// 	for _, v := range s {
// 		if v == str {
// 			return true
// 		}
// 	}

// 	return false
// }

// func ServerHandler(w http.ResponseWriter, r *http.Request) {
// 	ArtistsData = getdata.Binddata(baseUrl + "/artists")
// 	if ArtistsData == nil {
// 		fmt.Println("Error getting data")
// 		return
// 	}
// 	t, err := template.ParseFiles("templates/index.html")
// 	if err != nil {
// 		log.Fatalln(err)
// 		fmt.Fprint(w, "500 - Interal Server Error")
// 		return
// 	}
// 	if r.URL.Path != "/" {
// 		fmt.Fprint(w, "400 - Page not found")
// 		return
// 	}
// 	data := []getdata.FullData{}
// 	for _, artist := range ArtistsData {
// 		if contains(ArtistsData, artist.Name) {

// 		}
// 	}
// 	t.Execute(w, ArtistsData)
// }

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	link := "https://groupietrackers.herokuapp.com/api/artists"

	t, err := template.ParseFiles("templates/index_start.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, nil)

	var data []getdata.FullData
	data = getdata.BindData(link)
	for i, _ := range data {
		artT, err := template.ParseFiles("templates/artistThumb.html")
		if err != nil {
			log.Fatalln(err)
		}
		artT.Execute(w, data[i])
	}

	t, err = template.ParseFiles("templates/index_end.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, nil)
}
