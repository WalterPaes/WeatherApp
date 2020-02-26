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

type Data struct {
	City    []City
	Weather []weather.Weather
	Error   error
	svc     Service
}

func NewLocation(svc Service) *Data {
	return &Data{svc: svc}
}

func (l *Data) GetWeatherConditionsByCity(city string) {
	err := l.getLocationData(city)
	if err != nil {
		l.Error = err
		return
	}

	err = l.getCurrentCondition()
	if err != nil {
		return
	}

	return
}

func (l *Data) getLocationData(city string) error {
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

func (l *Data) getCurrentCondition() error {
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
