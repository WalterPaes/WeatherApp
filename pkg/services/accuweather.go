package services

import (
	"errors"
	"fmt"
	"github.com/WalterPaes/WeatherApp/pkg/lib/api"
	"io/ioutil"
	"net/http"
)

var language = "pt-br"
var errorsDetails = map[int]string{
	400: "Request had bad syntax or the parameters supplied were invalid",
	401: "Unauthorized. API authorization failed",
	403: "Unauthorized. You do not have permission to access this endpoint",
	404: "Server has not found a route matching the given URI",
	500: "Server encountered an unexpected condition which prevented it from fulfilling the request",
	503: "Service is unavaiable",
}

type Accuweather struct {
	url string
	key string
}

func NewAccuweather() *Accuweather {
	return &Accuweather{
		url: "http://dataservice.accuweather.com/locations/v1/cities/search",
		key: "j7wam502gD85s48G2POGGyzVb6sIGXWg",
	}
}

func (a Accuweather) GetLocationData(city string) *api.HttpResponse {
	url := "http://dataservice.accuweather.com/locations/v1/cities/search"
	params := map[string]string{
		"apikey":   a.key,
		"q":        city,
		"language": language,
		"details":  "true",
	}

	err := a.setUrl(url, params)
	if err != nil {
		return api.NewHttpResponse(nil, http.StatusInternalServerError, err)
	}

	response := a.doRequest()
	if response.HasError() {
		return api.NewHttpResponse(nil, response.HttpStatus, response.Error)
	}

	return api.NewHttpResponse(response.Body, response.HttpStatus, response.Error)
}

func (a Accuweather) GetCurrentCondition(locationKey string) *api.HttpResponse {
	url := fmt.Sprintf("http://dataservice.accuweather.com/currentconditions/v1/%s", locationKey)
	params := map[string]string{
		"apikey":   a.key,
		"language": language,
	}

	err := a.setUrl(url, params)
	if err != nil {
		return api.NewHttpResponse(nil, http.StatusInternalServerError, err)
	}

	response := a.doRequest()
	if response.HasError() {
		return api.NewHttpResponse(nil, response.HttpStatus, response.Error)
	}

	return api.NewHttpResponse(response.Body, response.HttpStatus, response.Error)
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

func (a Accuweather) doRequest() *api.HttpResponse {
	resp, err := http.Get(a.url)
	if err != nil {
		return api.NewHttpResponse(nil, http.StatusInternalServerError, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return api.NewHttpResponse(nil, resp.StatusCode, errors.New(errorsDetails[resp.StatusCode]))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return api.NewHttpResponse(nil, http.StatusInternalServerError, err)
	}

	return api.NewHttpResponse(body, resp.StatusCode, err)
}
