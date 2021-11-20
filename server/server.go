package main

import (
	"log"
	"net/http"
)

func (a *App) start() {
	log.Println("Server is Running on 0.0.0.0:8080")
	var film Film
	a.r.HandleFunc("/films", film.GetAllFilmsJson).Methods("GET")
	// a.r.HandleFunc("/students", a.addStudent).Methods("POST")
	// a.r.HandleFunc("/students/{id}", a.updateStudent).Methods("PUT")
	// a.r.HandleFunc("/students/{id}", a.deleteStudent).Methods("DELETE")
	a.r.PathPrefix("/").Handler(http.FileServer(http.Dir("./webapp/dist/webapp/")))
	log.Fatal(http.ListenAndServe(":8080", a.r))
}
