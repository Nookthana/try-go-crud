package main

import (
	"log"
	"net/http"

	s "golang/servies"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/mem", s.GetDataPersonID).Methods("GET")
	router.HandleFunc("/mem/{id}", s.GetDataPersonAll).Methods("GET")
	router.HandleFunc("/mem", s.CreatePerson).Methods("POST")
	router.HandleFunc("/mem/{id}", s.UpDatePersonID).Methods("PUT")
	router.HandleFunc("/mem/{id}", s.DeletePersonID).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9999", router))

}
