package binddata

type FullData struct {
	Id             int                 `json:"id"`
	Name           string              `json:"name"`
	Image          string              `json:"image"`
	Members        []string            `json:"members"`
	CreationDate   string              `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	Locations      []string            `json:"locations"`
	ConcertDates   []string            `json:"concerts"`
	DatesLocations map[string][]string `json:"datesLocations"`
}


type Artist struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate string   `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concerts"`
	Relations    string   `json:"relations"`
}

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}
type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Locations struct {
	Index []Location `json:"index"`
}

type Dates struct {
	Index []Date `json:"index"`
}

type Relations struct {
	Index []Relation `json:"index"`
}

