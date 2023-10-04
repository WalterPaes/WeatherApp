package main

import (
	"github.com/WalterPaes/WeatherApp/configs"
	"github.com/WalterPaes/WeatherApp/internal/weather/handlers"
	"github.com/WalterPaes/WeatherApp/internal/weather/services"
	accuweather2 "github.com/WalterPaes/WeatherApp/pkg/accuweather"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	configVars := configs.Load()

	accuweather := accuweather2.NewAccuweather(http.DefaultClient, configVars.AccuweatherBaseUrl, configVars.AccuweatherApiKey)
	weatherService := services.NewService(accuweather)
	weatherHandler := handlers.NewHandler(weatherService)

	r := mux.NewRouter()
	r.HandleFunc("/", weatherHandler.Index).Methods(http.MethodGet)
	r.HandleFunc("/conditions", weatherHandler.GetCurrentCondition).Methods(http.MethodPost)

	log.Println("Start Server...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err.Error())
	}
}
