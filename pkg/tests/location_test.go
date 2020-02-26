package tests

import (
	"errors"
	"github.com/WalterPaes/WeatherApp/pkg/domains/location"
	"github.com/WalterPaes/WeatherApp/pkg/lib/api"
	"net/http"
	"testing"
)

var MockError = errors.New("Ops...")

type MockService struct{}

func (ms *MockService) GetLocationData(city string) *api.HttpResponse {
	if city == "Florianópolis" {
		jsonPayload := `[{
			"LocalizedName": "Florianópolis",
			"Region": {"LocalizedName": "América do Sul"},
			"Country": {"LocalizedName": "Brasil"},
			"AdministrativeArea": {"LocalizedName": "Santa Catarina"},
			"GeoPosition": {"Latitude": -27.588,"Longitude": -48.548},
			"CanonicalLocationKey": "35952"
		}]`
		return api.NewHttpResponse([]byte(jsonPayload), http.StatusOK, nil)
	}

	return api.NewHttpResponse(nil, http.StatusInternalServerError, MockError)
}

func (ms *MockService) GetCurrentCondition(locationKey string) *api.HttpResponse {
	if locationKey == "35952" {
		jsonPayload := `[{
    		"WeatherText": "Uma pancada de chuva",
    		"WeatherIcon": 12,
    		"HasPrecipitation": true,
    		"Temperature": {"Metric": {"Value": 24.7,"Unit": "C"}}
		}]`
		return api.NewHttpResponse([]byte(jsonPayload), http.StatusOK, nil)
	}

	return api.NewHttpResponse(nil, http.StatusInternalServerError, MockError)
}

func TestLocation(t *testing.T) {
	t.Run("Success: Testing City Name", func(t *testing.T) {
		city := "Florianópolis"
		l := location.NewLocation(&MockService{})
		l.GetWeatherConditionsByCity(city)

		got := l.City[0].Name
		want := city

		if got != want {
			t.Errorf("Want '%v', but Got '%v'", want, got)
		}
	})

	t.Run("Error: Testing City Name", func(t *testing.T) {
		city := "Belém"
		l := location.NewLocation(&MockService{})
		l.GetWeatherConditionsByCity(city)

		got := l.Error
		want := MockError

		if got != want {
			t.Errorf("Want '%v', but Got '%v'", want, got)
		}
	})
}

func TestCondition(t *testing.T) {
	t.Run("Success: Correct Condition", func(t *testing.T) {
		city := "Florianópolis"
		l := location.NewLocation(&MockService{})
		l.GetWeatherConditionsByCity(city)

		got := l.Weather[0].Text
		want := "Uma pancada de chuva"

		if got != want {
			t.Errorf("Want '%v', but Got '%v'", want, got)
		}
	})
}