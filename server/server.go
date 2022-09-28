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

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	// checking if the path is not correct and returning 400
	if r.URL.Path != "/" && strings.Index(r.URL.Path, "/artists/") != 0 {
		fmt.Fprint(w, "400 - Page not found")
		return
	}
	// parsing the html file
	t, err := template.ParseFiles("templates/index_start.html")
	if err != nil {
		fmt.Fprint(w, "500 - Interal Server Error")
		return
	}
	t.Execute(w, nil)

	// getting the data from the api
	link := "https://groupietrackers.herokuapp.com/api/artists"

	data := getdata.BindData(link)
	if data == nil {
		fmt.Println("Error getting data")
		return
	}
	// getting the artist number from the url
	if strings.Contains(r.URL.Path, "/artists/")  {
		artistName := strings.Split(r.URL.Path, "/artists/")[1]
		id, _ := strconv.Atoi(artistName)
		artist := getdata.GetArtistById(link, id)
		art, err := template.ParseFiles("templates/artist.html")
		if err != nil {
			log.Fatalln(err)
		}
		art.Execute(w, artist)
	}
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
