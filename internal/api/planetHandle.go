package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"projeto-star-wars-api-go/internal/api/request"
	"projeto-star-wars-api-go/internal/planet"

	"github.com/gorilla/mux"
)

type PlanetHandler struct {
	service planet.Service
}

func NewPlanetHandler(service planet.Service) *PlanetHandler {
	return &PlanetHandler{service: service}
}
func (p *PlanetHandler) SavePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	var in request.PlanetIn
	err = json.Unmarshal(body, &in)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}

	document := in.ToModel()
	hexId, err := p.service.Save(context.Background(), document)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	url := "http://localhost:8080/planets/" + hexId
	w.Header().Add("location", url)
	w.WriteHeader(http.StatusCreated)

}
func (p *PlanetHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	planets, err := p.service.FindAll(context.Background())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(planets)
	w.WriteHeader(http.StatusOK)

}

func (p *PlanetHandler) FindById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	planet, err := p.service.FindById(context.Background(), vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//var in *response.PlanetOut
	//planet2 := in.FromModel(*planet)

	encoder := json.NewEncoder(w)
	encoder.Encode(planet)
	w.WriteHeader(http.StatusOK)

}

func (p *PlanetHandler) FindByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	planetName := r.URL.Query().Get("name")
	planet, err := p.service.FindByName(context.Background(), planetName)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(planet) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//var in *response.PlanetOut
	//var resp []*response.PlanetOut

	//for _, plan := range planet {

	//	plan2 := in.FromModel(plan)
	//	resp = append(resp, plan2) }

	encoder := json.NewEncoder(w)
	encoder.Encode(planet)
	w.WriteHeader(http.StatusOK)

}
func (p *PlanetHandler) UpdateById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var newPlanet request.PlanetIn
	err := json.NewDecoder(r.Body).Decode(&newPlanet)

	if err != nil {
		log.Println("Error Decoding the planet", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	planetModel := newPlanet.ToModel()

	planet, err := p.service.UpdateById(context.Background(), *planetModel, vars["id"])
	if err != nil {
		log.Println("Error updating the planet", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := json.Marshal(planet)

	if err != nil {
		log.Println("Error Marshaling the result planet", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		log.Println("Error to write the response", err)
	}

}

func (p *PlanetHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	err := p.service.DeleteById(context.Background(), vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
}
func (p *PlanetHandler) Healthcheck(w http.ResponseWriter, r *http.Request) {

	_, err := p.service.Healthcheck()

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	w.WriteHeader(http.StatusOK)
	return
}
