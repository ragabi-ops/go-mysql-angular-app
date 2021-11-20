package main

import "github.com/gorilla/mux"

func main() {
	app := App{
		r: mux.NewRouter(),
	}
	app.start()
}
