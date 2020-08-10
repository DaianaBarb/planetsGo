package main

import (

	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "os"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
)
//------------classe planeta
type Planet struct {
	ID                          string `json:"id"`
	Name                        string `json:"name"`
	Climate                     string `json:"climate"`
	Terrain                     string `json:"terrain"`
	NumberOfAppearancesInMovies int    `json:"numberOfAppearancesInMovies"`
}
type Planet2 struct {
_ID                          primitive.ObjectID`json:"_id"`
Name                        string `json:"name"`
Climate                     string `json:"climate"`
Terrain                     string `json:"terrain"`
NumberOfAppearancesInMovies int    `json:"numberOfAppearancesInMovies"`
}

//------------------ lista de planetas

var planets []Planet = []Planet{
	Planet{
		ID:                          "0",
		Name:                        "Tatooine",
		Climate:                     "arid ",
		Terrain:                     " solid ",
		NumberOfAppearancesInMovies: 5,
	},

	Planet{
		ID:                          "1",
		Name:                        "Tatooine",
		Climate:                     "arid ",
		Terrain:                     " solid ",
		NumberOfAppearancesInMovies: 5,
	},
	Planet{
		ID:                          "2",
		Name:                        "Tatooine",
		Climate:                     "arid ",
		Terrain:                     " solid ",
		NumberOfAppearancesInMovies: 5,
	},
	Planet{
		ID:                          "3",
		Name:                        "Tatooine",
		Climate:                     "arid ",
		Terrain:                     " solid ",
		NumberOfAppearancesInMovies: 5,
	},
}
//-------------funções dos endPoints
func savePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		//tratar erro

	}
	var newPlanet Planet2

	json.Unmarshal(body, &newPlanet)
	//newPlanet._ID=primitive.NewObjectID()
	collection.InsertOne(ctx,newPlanet)
	// planets = append(planets, newPlanet)

var	plan Planet

	plan.ID=newPlanet._ID.Hex()
	plan.Terrain=newPlanet.Terrain
	plan.Climate=newPlanet.Climate
	plan.Name=newPlanet.Name
	plan.NumberOfAppearancesInMovies=newPlanet.NumberOfAppearancesInMovies
	encoder := json.NewEncoder(w)
	encoder.Encode(plan)
	w.WriteHeader(http.StatusCreated)

}
func returnPlanetId(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	vars := mux.Vars(r)
	id:= vars["id"]
	for _,planet := range planets{
		if planet.ID==id{
			json.NewEncoder(w).Encode(planet)
		}
	}
   w.WriteHeader(http.StatusNotFound)
}
func returnPlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(planets)

}


func deletePlanet(w http.ResponseWriter, r *http.Request) {

	vars:= mux.Vars(r)
	id:= vars["id"]
	var plan Planet
	planVazia :=plan

	for _,planet := range planets{
		if planet.ID==id{
			plan=planet
		}
	}
	if(plan == planVazia){
		w.WriteHeader(http.StatusNotFound)
		return 
	}

	idd, _ :=strconv.Atoi(id)
	planets = append(planets[:idd], planets[idd+1:]...)
	w.WriteHeader(http.StatusOK)

}

func updatePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		//tratar erro

	}
	var newPlanet Planet
	json.Unmarshal(body, &newPlanet)
	vars := mux.Vars(r)
	id:= vars["id"]
	indice := ""
	indice ="null"
	for _,planet := range planets{
		if planet.ID==id {
	       indice = id
  break


}	}
	if indice=="null"{
	w.WriteHeader(http.StatusNotFound)
	return}
	idInt,_:= strconv.Atoi(id)
	planets[idInt]=newPlanet
	w.WriteHeader(http.StatusOK)
}
//-----------banco de dados
var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("tasker").Collection("tasks")
}


func main() {
	fmt.Println("Servidor esta rodando na porta 8080")
	router := mux.NewRouter()
	router.HandleFunc("/delete/{id}", deletePlanet).Methods("DELETE")
	router.HandleFunc("/update/{id}", updatePlanet).Methods("PUT")
	router.HandleFunc("/save", savePlanet).Methods("POST")
	router.HandleFunc("/planets", returnPlanet).Methods("GET")
	router.HandleFunc("/planet/{id}", returnPlanetId).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}



