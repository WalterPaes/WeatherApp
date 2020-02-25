package weather

import (
	"github.com/WalterPaes/WeatherApp/pkg/domains/location"
	"github.com/WalterPaes/WeatherApp/pkg/services"
	"log"
	"net/http"
)

var service = services.NewAccuweather()

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	l := location.NewLocation(service)
	err := l.GetWeatherConditionsByCity(r.FormValue("city"))
	if err != nil {
		log.Println(err.Error())
	}
}
