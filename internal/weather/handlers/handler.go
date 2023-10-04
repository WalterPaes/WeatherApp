package handlers

import (
	"github.com/WalterPaes/WeatherApp/internal/weather"
	"html/template"
	"log"
	"net/http"
)

type Handler struct {
	service weather.WeatherService
}

type data struct {
	Data  any
	Error error
}

func NewHandler(service weather.WeatherService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderTemplate(w, nil)
		return
	}
}

func (h *Handler) GetCurrentCondition(w http.ResponseWriter, r *http.Request) {
	city := r.FormValue("city")
	currentCondition, err := h.service.GetConditions(city)

	renderTemplate(w, &data{
		Data:  currentCondition,
		Error: err,
	})
}

func renderTemplate(w http.ResponseWriter, d *data) {
	tpl := template.Must(template.ParseFiles("pkg/ui/layout.html"))
	err := tpl.Execute(w, d)
	if err != nil {
		log.Println(err.Error())
	}
}
