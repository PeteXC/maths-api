package main

import (
	"net/http"

	"github.com/PeteXC/maths-api/api/handlers/maths/add"
	"github.com/gorilla/mux"
	"github.com/joerdav/zapray"
)

func main() {
	//things to go in handler
	log, err := zapray.NewProduction()
	if err != nil {
		panic("failed to create logger: " + err.Error())
	}

	//handler
	addHandler, err := add.NewHandler(log)

	r := mux.NewRouter()
	if err != nil {
		panic("I can't create a router")
	}

	r.Handle("/maths/add", addHandler)

	log.Info("Hey I've started the local api")

	err = http.ListenAndServe("127.0.0.1:8080", r)

}
