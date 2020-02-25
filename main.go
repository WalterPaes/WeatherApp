package main

import (
	"github.com/WalterPaes/WeatherApp/pkg/domains/location"
	"github.com/WalterPaes/WeatherApp/pkg/services"
)

var service = services.NewAccuweather()

func main()  {
	l := location.NewLocation(service)
	err := l.GetWeatherConditionsByCity("Belém")
	if err != nil {
		panic(err)
	}
}
