package api

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"projeto-star-wars-api-go/internal/planet"
)

type PlanetHandler struct {
	service *planet.Service
}

func NewPlanetHandler(service *planet.Service) *PlanetHandler {
	return &PlanetHandler{service: service}
}
func (p *PlanetHandler) SavePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		//tratar erro
	}
	var in planet.PlanetIn
	error = json.Unmarshal(body, &in)
	if error != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}

	document := in.ToDocument()
	error = p.service.Save(context.Background(), document)
	url := "http://localhost:8080/planets/" + document.ID.Hex()
	w.Header().Add("location", url)
	w.WriteHeader(http.StatusCreated)

}
func (p *PlanetHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	 planets, err := p.service.FindAll(context.Background())

  if err!=nil{
  	w.WriteHeader(http.StatusInternalServerError)
   }

	encoder := json.NewEncoder(w)
	encoder.Encode(planets)
	w.WriteHeader(http.StatusOK)

}

func (p *PlanetHandler) FindById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	planet, err := p.service.FindById(context.Background(),vars["id"])

	if err!=nil{
		w.WriteHeader(http.StatusNotFound)
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(planet)
	w.WriteHeader(http.StatusOK)

}

func (p *PlanetHandler) FindByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    vars := mux.Vars(r)
	planet, err := p.service.FindByName(context.Background(),vars["name"])

	  if err!=nil{
		w.WriteHeader(http.StatusNotFound)
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(planet)
	w.WriteHeader(http.StatusOK)

}
func (p *PlanetHandler) UpdateById(w http.ResponseWriter, r *http.Request) {

       vars := mux.Vars(r)
	   var newPlanet planet.PlanetIn
       err := json.NewDecoder(r.Body).Decode(&newPlanet)

	         		if err != nil {
	         			log.Println("Error Decoding the planet", err)
	 	                w.WriteHeader(http.StatusBadRequest)
						return
		}

		planet, err := p.service.UpdateById(context.Background(),newPlanet,vars["id"])
	               if err != nil{
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
	err := p.service.DeleteById(context.Background(),vars["id"])
	   if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
}

