package main

import (
	"github.com/PeteXC/maths-api/api/handlers/maths/add"
	"github.com/akrylysov/algnhsa"
	"github.com/joerdav/zapray"
)

func main() {

	//things to go in handler
	log, err := zapray.NewProduction()
	if err != nil {
		panic("failed to create logger: " + err.Error())
	}

	//handler
	handler, err := add.NewHandler(log)

	xhandler := zapray.NewMiddleware("mathsAdPostHandler", handler)

	//algnhsa.listenandserve
	algnhsa.ListenAndServe(xhandler, nil)
}
