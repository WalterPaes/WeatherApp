package location

import (
	"encoding/json"
	"github.com/WalterPaes/WeatherApp/pkg/domains/weather"
	"github.com/WalterPaes/WeatherApp/pkg/lib/api"
)

type Service interface {
	GetLocationData(city string) *api.HttpResponse
	GetCurrentCondition(locationKey string) *api.HttpResponse
}

type Location struct {
	City    []City
	Weather []weather.Weather
	svc     Service
}

func NewLocation(svc Service) *Location {
	return &Location{svc: svc}
}

func (l *Location) GetWeatherConditionsByCity(city string) error {
	err := l.getLocationData(city)
	if err != nil {
		return err
	}

	err = l.getCurrentCondition()
	if err != nil {
		return err
	}

	return nil
}

func (l *Location) getLocationData(city string) error {
	response := l.svc.GetLocationData(city)
	if response.HasError() {
		return response.Error
	}

	err := json.Unmarshal(response.Body, &l.City)
	if err != nil {
		return err
	}

	return nil
}

func (l *Location) getCurrentCondition() error {
	response := l.svc.GetCurrentCondition(l.City[0].Details.LocationKey)
	if response.HasError() {
		return response.Error
	}

	err := json.Unmarshal(response.Body, &l.Weather)
	if err != nil {
		return err
	}

	return nil
}
