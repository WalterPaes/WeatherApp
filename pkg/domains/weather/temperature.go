package weather

type Temperature struct {
	Metric struct {
		Value float32 `json:"Value"`
		Unit  string  `json:"Unit"`
	} `json:"Metric"`
}
