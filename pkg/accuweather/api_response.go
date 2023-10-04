package accuweather

type SearchCitiesResponse struct {
	Cities []Cities
}

type Cities struct {
	Version           int    `json:"Version"`
	Key               string `json:"Key"`
	Type              string `json:"Type"`
	Rank              int    `json:"Rank"`
	LocalizedName     string `json:"LocalizedName"`
	EnglishName       string `json:"EnglishName"`
	PrimaryPostalCode string `json:"PrimaryPostalCode"`
	Region            struct {
		ID            string `json:"ID"`
		LocalizedName string `json:"LocalizedName"`
		EnglishName   string `json:"EnglishName"`
	} `json:"Region"`
	Country struct {
		ID            string `json:"ID"`
		LocalizedName string `json:"LocalizedName"`
		EnglishName   string `json:"EnglishName"`
	} `json:"Country"`
	AdministrativeArea struct {
		ID            string `json:"ID"`
		LocalizedName string `json:"LocalizedName"`
		EnglishName   string `json:"EnglishName"`
		Level         int    `json:"Level"`
		LocalizedType string `json:"LocalizedType"`
		EnglishType   string `json:"EnglishType"`
		CountryID     string `json:"CountryID"`
	} `json:"AdministrativeArea"`
	TimeZone struct {
		Code             string  `json:"Code"`
		Name             string  `json:"Name"`
		GmtOffset        float32 `json:"GmtOffset"`
		IsDaylightSaving bool    `json:"IsDaylightSaving"`
		NextOffsetChange any     `json:"NextOffsetChange"`
	} `json:"TimeZone"`
	GeoPosition struct {
		Latitude  float64 `json:"Latitude"`
		Longitude float64 `json:"Longitude"`
		Elevation struct {
			Metric struct {
				Value    float32 `json:"Value"`
				Unit     string  `json:"Unit"`
				UnitType int     `json:"UnitType"`
			} `json:"Metric"`
			Imperial struct {
				Value    float32 `json:"Value"`
				Unit     string  `json:"Unit"`
				UnitType int     `json:"UnitType"`
			} `json:"Imperial"`
		} `json:"Elevation"`
	} `json:"GeoPosition"`
	IsAlias    bool `json:"IsAlias"`
	ParentCity struct {
		Key           string `json:"Key"`
		LocalizedName string `json:"LocalizedName"`
		EnglishName   string `json:"EnglishName"`
	} `json:"ParentCity"`
	SupplementalAdminAreas []struct {
		Level         int    `json:"Level"`
		LocalizedName string `json:"LocalizedName"`
		EnglishName   string `json:"EnglishName"`
	} `json:"SupplementalAdminAreas"`
	DataSets []string `json:"DataSets"`
}

type GetConditionsResponse struct {
	Conditions []Conditions
}

type Conditions struct {
	LocalObservationDateTime string `json:"LocalObservationDateTime"`
	EpochTime                int    `json:"EpochTime"`
	WeatherText              string `json:"WeatherText"`
	WeatherIcon              int    `json:"WeatherIcon"`
	HasPrecipitation         bool   `json:"HasPrecipitation"`
	PrecipitationType        any    `json:"PrecipitationType"`
	IsDayTime                bool   `json:"IsDayTime"`
	Temperature              struct {
		Metric struct {
			Value    float64 `json:"Value"`
			Unit     string  `json:"Unit"`
			UnitType int     `json:"UnitType"`
		} `json:"Metric"`
		Imperial struct {
			Value    float32 `json:"Value"`
			Unit     string  `json:"Unit"`
			UnitType int     `json:"UnitType"`
		} `json:"Imperial"`
	} `json:"Temperature"`
	MobileLink string `json:"MobileLink"`
	Link       string `json:"Link"`
}
