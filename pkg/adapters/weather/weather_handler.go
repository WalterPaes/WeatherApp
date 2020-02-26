package weather

import (
	"github.com/WalterPaes/WeatherApp/pkg/domains/location"
	"github.com/WalterPaes/WeatherApp/pkg/services"
	"html/template"
	"log"
	"net/http"
)

var service = services.NewAccuweather()

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderTemplate(w, nil)
		return
	}

	l := location.NewLocation(service)
	l.GetWeatherConditionsByCity(r.FormValue("city"))
	if l.Error != nil {
		log.Println(l.Error.Error())
	}
	renderTemplate(w, l)
}

func renderTemplate(w http.ResponseWriter, data *location.Data) {
	tpl := template.Must(template.ParseFiles("pkg/ui/layout.html"))
	err := tpl.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
	}
}
