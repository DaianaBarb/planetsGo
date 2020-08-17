package main

import (
	"fmt"
	"log"
	"net/http"
	"projeto-star-wars-api-go/internal/api"
	"projeto-star-wars-api-go/internal/planet"
	dao2 "projeto-star-wars-api-go/internal/provider/mongo/dao"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Servidor esta rodando na porta 8080")
	database, _ := dao2.GetDatabase()
	dao := dao2.NewMongoPlanet(database)
	service := planet.NewService(dao)
	handler := api.NewPlanetHandler(service)
	router := mux.NewRouter()
	router.HandleFunc("/planets/{id}", handler.DeleteById).Methods("DELETE")
	router.HandleFunc("/planets/{id}", handler.UpdateById).Methods("PUT")
	router.HandleFunc("/planets", handler.SavePlanet).Methods("POST")
	router.HandleFunc("/planets", handler.GetAll).Methods("GET")
	router.HandleFunc("/planets/{id}", handler.FindById).Methods("GET")
	router.HandleFunc("/planets/", handler.FindByName).Methods("GET").Queries("name", "")
	router.HandleFunc("planets/healthcheck", handler.Healthcheck).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}

//func getDatabase() *mongo.Database {
//clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
//	client, err := mongo.Connect(context.Background(), clientOptions)
//	if err != nil {
//tratar
//}
//return client.Database("star-wars")
//}
