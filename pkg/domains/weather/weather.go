package weather

type Weather struct {
	Text          string      `json:"WeatherText"`
	Icon          int         `json:"WeatherIcon"`
	Precipitation bool        `json:"HasPrecipitation"`
	Temperature   Temperature `json:"Temperature"`
}
