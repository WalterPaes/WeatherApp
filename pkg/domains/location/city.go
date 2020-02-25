package location

type City struct {
	Name        string      `json:"LocalizedName"`
	Region      Region      `json:"Region"`
	Country     Country     `json:"Country"`
	State       State       `json:"AdministrativeArea"`
	GeoPosition GeoPosition `json:"GeoPosition"`
	Details     Details     `json:"Details"`
}
