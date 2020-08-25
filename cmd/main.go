package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"projeto-star-wars-api-go/internal/provider/mongo/dao"
	"projeto-star-wars-api-go/internal/router"
	"projeto-star-wars-api-go/internal/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Servidor esta rodando na porta 8080")

	database, err := getDatabase()
	if err != nil {
		log.Fatal("error get database")
	}

	dao := dao.NewMongoPlanet(database)
	swapi := service.NewSWAPI()
	service := service.NewPlanet(dao, swapi)
	handler := router.NewPlanetHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/planets/{id}", handler.DeleteById).Methods("DELETE")
	router.HandleFunc("/planets/{id}", handler.Update).Methods("PUT")
	router.HandleFunc("/planets", handler.SavePlanet).Methods("POST")
	router.HandleFunc("/planets", handler.FindAll).Methods("GET")
	router.HandleFunc("/planets/{id}", handler.FindById).Methods("GET")
	router.HandleFunc("/planets/", handler.FindByName).Methods("GET").Queries("name", "")
	router.HandleFunc("/health", handler.Healthcheck).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getDatabase() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	return client.Database("star-wars"), nil
}
