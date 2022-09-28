package server

import (
	"fmt"
	"gp/getdata"
	"log"
	"net/http"
	"strconv"
	"strings"
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

// getting the data from the api
var link = "https://groupietrackers.herokuapp.com/api/artists"

var data = getdata.BindData(link)

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	// checking if the path is not correct and returning 400
	if r.URL.Path != "/" && strings.Index(r.URL.Path, "/artists/") != 0 {
		fmt.Fprint(w, "400 - Page not found")
		return
	}
	// checking if the data is loaded correctly
	if data == nil {
		fmt.Println("Error getting data")
		return
	}
	// parsing the html file
	t, err := template.ParseFiles("templates/index_start.html")
	if err != nil {
		fmt.Fprint(w, "500 - Interal Server Error")
		return
	}
	t.Execute(w, nil)

	// getting the artist number from the url
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

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "/artists/") {
		// getting the artist number from the url
		artistId := strings.TrimPrefix(r.URL.Path, "/artists/")
		id, _ := strconv.Atoi(artistId)
		// getting the artist data
		artist := getdata.GetArtistById(link, id)
		// parsing the html file
		t, err := template.ParseFiles("templates/artist.html")
		if err != nil {
			log.Fatalln(err)
		}
		t.Execute(w, artist)
	}
}