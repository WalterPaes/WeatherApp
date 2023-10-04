package services

import "github.com/WalterPaes/WeatherApp/internal/weather"

type Service struct {
	integration weather.ApiIntegration
}

func NewService(integration weather.ApiIntegration) *Service {
	return &Service{integration: integration}
}

func (s Service) GetConditions(term string) (*weather.CurrentCondition, error) {
	cities, err := s.integration.CitiesSearch(term)
	if err != nil {
		return nil, err
	}

	city := cities.Cities[0]

	currentCondition, err := s.integration.CurrentCondition(city.Key)
	if err != nil {
		return nil, err
	}

	condition := currentCondition.Conditions[0]

	return &weather.CurrentCondition{
		City: weather.City{
			Name:    city.LocalizedName,
			Region:  city.Region.LocalizedName,
			Country: city.Country.LocalizedName,
			State:   city.AdministrativeArea.LocalizedName,
			GeoPosition: weather.GeoPosition{
				Latitude:  city.GeoPosition.Latitude,
				Longitude: city.GeoPosition.Longitude,
			},
		},
		Weather: weather.Weather{
			Text:          condition.WeatherText,
			IconId:        condition.WeatherIcon,
			Precipitation: condition.HasPrecipitation,
			Temperature: weather.Temperature{
				Value: condition.Temperature.Metric.Value,
			},
		},
	}, nil
}
