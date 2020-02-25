package main

import (
	"github.com/WalterPaes/WeatherApp/pkg/adapters/weather"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", weather.Handler).Methods(http.MethodGet)
	r.HandleFunc("/conditions", weather.Handler).Methods(http.MethodPost)

	log.Println("Start Server...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err.Error())
	}
}
