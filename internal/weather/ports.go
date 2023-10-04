package weather

import "github.com/WalterPaes/WeatherApp/pkg/accuweather"

type WeatherService interface {
	GetConditions(term string) (*CurrentCondition, error)
}

type ApiIntegration interface {
	CitiesSearch(city string) (*accuweather.SearchCitiesResponse, error)
	CurrentCondition(locationKey string) (*accuweather.GetConditionsResponse, error)
}
