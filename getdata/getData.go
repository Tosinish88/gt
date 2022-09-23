package getdata

import (
	"encoding/json"
	"fmt"
	"gp/binddata"
	"io/ioutil"
	"log"
	"net/http"
)

const baseUrl = "https://groupietrackers.herokuapp.com/api"

var Artists []binddata.Artist
var ArtistLocation binddata.Locations
var ArtistDates binddata.Dates
var ArtistRelations binddata.Relations
var ArtistData []binddata.ArtistData

func GetArtistData() {
	r, err := http.Get(baseUrl + "/artists")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &Artists)
	// fmt.Println(Artists)
}

func GetArtistLocation() {
	r, err := http.Get(baseUrl + "/locations")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &ArtistLocation)
	// fmt.Println(ArtistLocation)
}

func GetArtistDates() {
	r, err := http.Get(baseUrl + "/dates")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &ArtistDates)
	// fmt.Println(ArtistDates)
}

func GetArtistRelations() {
	r, err := http.Get(baseUrl + "/relation")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &ArtistRelations)
	// fmt.Println(ArtistRelations)
}

func GetData() []binddata.ArtistData {
	GetArtistData()
	GetArtistLocation()
	GetArtistDates()
	GetArtistRelations()

	for i := range Artists {
		var temp binddata.ArtistData
		temp.Id = i + 1
		temp.Name = Artists[i].Name
		temp.Image = Artists[i].Image
		temp.Members = Artists[i].Members
		temp.CreationDate = Artists[i].CreationDate
		temp.FirstAlbum = Artists[i].FirstAlbum
		temp.Locations = ArtistLocation.Index[i].Locations
		temp.ConcertDates = ArtistDates.Index[i].Dates
		temp.DatesLocations = ArtistRelations.Index[i].DatesLocations
		ArtistData = append(ArtistData, temp)
	}
	return ArtistData
}

func GetArtistById(id int) binddata.ArtistData {
	data := GetData()
	// fmt.Println(len(data))
	if id <= 0 || id > len(data) {
		fmt.Println("Invalid id")
		return binddata.ArtistData{}
	}
	// fmt.Println(GetData()[id-1])
	return data[id-1]
}
