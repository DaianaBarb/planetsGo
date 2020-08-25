package router

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"projeto-star-wars-api-go/internal/model"
	"projeto-star-wars-api-go/internal/service"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type PlanetHandler struct {
	service service.Planet
}

func NewPlanetHandler(service service.Planet) *PlanetHandler {
	return &PlanetHandler{service: service}
}
func (p *PlanetHandler) SavePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	var in model.PlanetIn
	err = json.Unmarshal(body, &in)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	error := ValidateStruct(&in)
	if error != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	hexId, err := p.service.Save(context.Background(), &in)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	url := "http://localhost:8080/planets/" + hexId
	w.Header().Add("location", url)
	w.WriteHeader(http.StatusCreated)

}
func (p *PlanetHandler) FindAll(w http.ResponseWriter, r *http.Request) {
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

	encoder := json.NewEncoder(w)
	encoder.Encode(planet)
	w.WriteHeader(http.StatusOK)
}

func (p *PlanetHandler) FindByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	planetName := r.URL.Query().Get("name")
	planets, err := p.service.FindByName(context.Background(), planetName)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(planets) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(planets)
	w.WriteHeader(http.StatusOK)

}
func (p *PlanetHandler) Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var planetIn model.PlanetIn
	json.NewDecoder(r.Body).Decode(&planetIn)

	//if err != nil {
	//log.Println("Error Decoding the planet", err)
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	error := ValidateStruct(&planetIn)
	if error != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	planet := p.service.Update(context.Background(), planetIn, vars["id"])

	if planet != nil {
		log.Println("Error updating the planet", planet)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *PlanetHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	err := p.service.DeleteById(context.Background(), vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func (p *PlanetHandler) Healthcheck(w http.ResponseWriter, r *http.Request) {

	error := GetDatabase()
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
func ValidateStruct(v *model.PlanetIn) error {
	var validate *validator.Validate
	validate = validator.New()

	errs := validate.Struct(v)
	if errs != nil {
		return errs
	}
	return nil
}

func GetDatabase() error {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	errr := client.Ping(ctx, nil)
	if errr != nil {
		return errr
	}
	return nil

}
