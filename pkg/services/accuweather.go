package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var language = "pt-br"

type Accuweather struct {
	url string
	key string
}

func NewAccuweather() *Accuweather {
	return &Accuweather{
		url: "http://dataservice.accuweather.com/locations/v1/cities/search",
		key: "DBo3uAAixdGNk8ph5Lzq1jgj75SDLxDV",
	}
}

func (a Accuweather) GetLocationData(city string) ([]byte, error) {
	url := "http://dataservice.accuweather.com/locations/v1/cities/search"
	params := map[string]string{
		"apikey": a.key,
		"q": city,
		"language": language,
		"details": "true",
	}

	err := a.setUrl(url, params)
	if err != nil {
		return nil, err
	}

	data, err := a.doRequest()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (a Accuweather) GetCurrentCondition(locationKey string) ([]byte, error) {
	url := fmt.Sprintf("http://dataservice.accuweather.com/currentconditions/v1/%s", locationKey)
	params := map[string]string{
		"apikey": a.key,
		"language": language,
	}

	err := a.setUrl(url, params)
	if err != nil {
		return nil, err
	}

	data, err := a.doRequest()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (a *Accuweather) setUrl(url string, params map[string]string) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	// Setting URLS params
	query := req.URL.Query()
	for i, v := range params {
		query.Add(i, v)
	}

	// URL Encode
	req.URL.RawQuery = query.Encode()

	// Set to Struct
	a.url = req.URL.String()

	return nil
}

func (a Accuweather) doRequest() ([]byte, error) {
	resp, err := http.Get(a.url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
