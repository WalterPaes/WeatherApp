package accuweather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var errorsDetails = map[int]string{
	http.StatusBadRequest:          "Request had bad syntax or the parameters supplied were invalid",
	http.StatusUnauthorized:        "Unauthorized. API authorization failed",
	http.StatusForbidden:           "Forbidden. You do not have permission to access this endpoint",
	http.StatusNotFound:            "Server has not found a route matching the given URI",
	http.StatusInternalServerError: "Server encountered an unexpected condition which prevented it from fulfilling the request",
	http.StatusServiceUnavailable:  "Service is unavailable",
}

type Accuweather struct {
	client  *http.Client
	baseUrl string
	apiKey  string
}

func NewAccuweather(client *http.Client, baseUrl, apiKey string) *Accuweather {
	return &Accuweather{
		client:  client,
		baseUrl: baseUrl,
		apiKey:  apiKey,
	}
}

func (a Accuweather) CitiesSearch(city string) (*SearchCitiesResponse, error) {
	url := fmt.Sprintf("%s/locations/v1/cities/search", a.baseUrl)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	params := map[string]string{
		"apikey": a.apiKey,
		"q":      city,
	}

	query := req.URL.Query()
	for i, v := range params {
		query.Add(i, v)
	}
	req.URL.RawQuery = query.Encode()

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(errorsDetails[resp.StatusCode])
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response []Cities
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &SearchCitiesResponse{Cities: response}, nil
}

func (a Accuweather) CurrentCondition(locationKey string) (*GetConditionsResponse, error) {
	url := fmt.Sprintf("%s/currentconditions/v1/%s", a.baseUrl, locationKey)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	params := map[string]string{
		"apikey":   a.apiKey,
		"language": "pt-br",
	}

	query := req.URL.Query()
	for i, v := range params {
		query.Add(i, v)
	}

	req.URL.RawQuery = query.Encode()

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(errorsDetails[resp.StatusCode])
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response []Conditions
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &GetConditionsResponse{Conditions: response}, nil
}
