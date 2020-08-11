package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"projeto-star-wars-api-go/internal/planet"
)

type PlanetHandler struct {
	service *planet.Service
}
func NewPlanetHandler(service *planet.Service) *PlanetHandler{
	return &PlanetHandler{service:service}
}
func (p *PlanetHandler)SavePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, error := ioutil.ReadAll(r.Body)
	if error !=nil{
 //tratar erro
	}
	var in planet.PlanetIn
  error= json.Unmarshal(body, &in)
  if error != nil {
  	w.WriteHeader(http.StatusUnprocessableEntity)
  }

  document:= in.ToDocument()
  error = p.service.Save(context.Background(),document)
  url := "http://localhost:8080/planets/"+ document.ID.Hex()
   w.Header().Add("location",url)
	w.WriteHeader(http.StatusCreated)


}



