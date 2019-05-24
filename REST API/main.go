package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"./arango"
	"github.com/gorilla/mux"
)

// The boton Type (more like an object)
type Boton struct {
	ID     string `json:"ID"`
	LUMINO string `json:"LUMINO"`

	TEMP    string `json:"TEMP"`
	HUMEDAD string `json:"HUMEDAD"`
	PH      string `json:"PH"`

	BOMBA string `json:"BOMBA"`
	UV    string `json:"UV"`
}

var botoncito []Boton

// Display all from thebotoncito var
func GetBoton(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(arango.ObtenerTodo())

}

// Display a single data
func GetBotoncito(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	/*for _, item := range botoncito {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}*/
	identificador := params["id"]
	json.NewEncoder(w).Encode(arango.Obtener(identificador))
}

// create a new item
func CreateBoton(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var boton arango.Boton
	_ = json.NewDecoder(r.Body).Decode(&boton)

	arango.Conexion(boton)

	fmt.Println(json.NewEncoder(w).Encode(boton))
	fmt.Println(params)

}

// Delete an item
func DeleteBoton(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range botoncito {
		if item.ID == params["id"] {
			botoncito = append(botoncito[:index], botoncito[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(botoncito)
	}
}

// main function to boot up everything
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/boton", GetBoton).Methods("GET")
	router.HandleFunc("/boton/{id}", GetBotoncito).Methods("GET")
	router.HandleFunc("/boton", CreateBoton).Methods("POST")
	log.Fatal(http.ListenAndServe(":800", router))

}
