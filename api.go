package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Status struct {
	St string `json:"status"`
}

var sttus Status

func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sttus.St = "ok"
	encoder := json.NewEncoder(w)
	encoder.Encode(sttus)

}

func configurarRotas() {
	http.HandleFunc("/", rotaPrincipal)

}

func configurarServidor() {
	configurarRotas()

	fmt.Println("Servidor esta rodando na porta 1337")
	log.Fatal(http.ListenAndServe(":1337", nil))

}

func main() {
	configurarServidor()
}
