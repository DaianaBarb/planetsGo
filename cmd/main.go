package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"projeto-star-wars-api-go/internal/api"
	"projeto-star-wars-api-go/internal/planet"
)

func main() {
	fmt.Println("Servidor esta rodando na porta 8080")
	database:= getDatabase()
	service:= planet.NewService(database)
     handler := api.NewPlanetHandler(service)
     router := mux.NewRouter()
	//router.HandleFunc("/delete/{id}", deletePlanet).Methods("DELETE")
	//router.HandleFunc("/update/{id}", updatePlanet).Methods("PUT")
	router.HandleFunc("/save", handler.SavePlanet).Methods("POST")
	//router.HandleFunc("/", handler.GetPlanets).Methods("GET")
	//router.HandleFunc("/planet/{id}", returnPlanetId).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func getDatabase() *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client,err :=mongo.Connect(context.Background(),clientOptions)
	if err !=nil{
		//tratar
	}
	return client.Database("star-wars")

}
