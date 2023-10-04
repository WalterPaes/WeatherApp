package weather

type CurrentCondition struct {
	City    `json:"city"`
	Weather `json:"weather"`
}

type City struct {
	Name        string      `json:"name"`
	Region      string      `json:"region"`
	Country     string      `json:"country"`
	State       string      `json:"state"`
	GeoPosition GeoPosition `json:"GeoPosition"`
}

type GeoPosition struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Weather struct {
	Text          string      `json:"text"`
	IconId        int         `json:"icon_id"`
	Precipitation bool        `json:"has_precipitation"`
	Temperature   Temperature `json:"temperature"`
}

type Temperature struct {
	Value float64 `json:"Value"`
}
