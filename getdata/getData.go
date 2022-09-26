package getdata

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type FullData struct {
	Id             int                 `json:"id"`
	Name           string              `json:"name"`
	Image          string              `json:"image"`
	Members        []string            `json:"members"`
	CreationDate   string              `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	Locations      []string            `json:"locations"`
	Relation	string            `json:"concerts"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func GetData(link string) []byte {
	r, err := http.Get(link)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return body
}

func Binddata(link string) []FullData {
	data := GetData(link)
	artists := []FullData{}

	err := json.Unmarshal(data, &artists)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	for i := 0; i < len(artists); i++ {
		r := Relation{}
		json.Unmarshal(GetData(artists[i].Relation), &r)
		artists[i].DatesLocations = r.DatesLocations
	}
	return artists
}


// func GetArtistLocation() {
// 	r, err := http.Get(baseUrl + "/locations")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	json.Unmarshal(body, &ArtistLocation)
// 	// fmt.Println(ArtistLocation)
// }

// func GetArtistDates() {
// 	r, err := http.Get(baseUrl + "/dates")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	json.Unmarshal(body, &ArtistDates)
// 	// fmt.Println(ArtistDates)
// }

// func GetArtistRelations() {
// 	r, err := http.Get(baseUrl + "/relation")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	json.Unmarshal(body, &ArtistRelations)
// 	// fmt.Println(ArtistRelations)
// }

// var ArtistsFullData []binddata.FullData

// func GetData() []binddata.FullData {
// 	if ArtistsFullData != nil {
// 		return ArtistsFullData
// 	}
// 	GetArtistData()
// 	GetArtistLocation()
// 	GetArtistDates()
// 	GetArtistRelations()

// 	for i := range Artists {
// 		var temp binddata.FullData
// 		temp.Id = i + 1
// 		temp.Name = Artists[i].Name
// 		temp.Image = Artists[i].Image
// 		temp.Members = Artists[i].Members
// 		temp.CreationDate = Artists[i].CreationDate
// 		temp.FirstAlbum = Artists[i].FirstAlbum
// 		temp.Locations = ArtistLocation.Index[i].Locations
// 		temp.ConcertDates = ArtistDates.Index[i].Dates
// 		temp.DatesLocations = ArtistRelations.Index[i].DatesLocations
// 		ArtistsFullData = append(ArtistsFullData, temp)
// 	}
// 	return ArtistsFullData
// }

// func GetArtistById(id int) binddata.FullData {
// 	for _, artist := range GetData() {
// 		if artist.Id == id {
// 			return artist
// 		}
// 	}
// 	return binddata.FullData{}
// }
