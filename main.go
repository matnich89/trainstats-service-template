package main

import (
	"github.com/go-chi/chi/v5"
	cmd "goapp-temmplate/cmd/api"
	"goapp-temmplate/handler"
	"log"
)

func main() {
	router := chi.NewMux()

	requestHandler := handler.NewHandler()

	app := cmd.NewApp(router, requestHandler)

	err := app.Serve()

	if err != nil {
		log.Fatalln(err)
	}

}
