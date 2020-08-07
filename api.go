package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Status struct {
	St string `json:"status"`
}

var status Status = Status{"ok"}
var sttus Status

func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sttus.St = "ok"
	encoder := json.NewEncoder(w)
	encoder.Encode(status)

}

func main() {
	fmt.Println("Servidor esta rodando na porta 8080")
	router := mux.NewRouter()

	router.HandleFunc("/status", rotaPrincipal).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}
