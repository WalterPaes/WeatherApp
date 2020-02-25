package location

import (
	"encoding/json"
	"github.com/WalterPaes/WeatherApp/pkg/domains/weather"
)

type Service interface {
	GetLocationData(city string) ([]byte, error)
	GetCurrentCondition(locationKey string) ([]byte, error)
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
	data, err := l.svc.GetLocationData(city)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &l.City)
	if err != nil {
		return err
	}

	return nil
}

func (l *Location) getCurrentCondition() error {
	data, err := l.svc.GetCurrentCondition(l.City[0].Details.LocationKey)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &l.Weather)
	if err != nil {
		return err
	}

	return nil
}
