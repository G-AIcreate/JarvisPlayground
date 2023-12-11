package main

import (
	"log"
	"net/http"
	// todo import problem needs to be fixed
	_ "github.com/go-chi/chi/v5"
	_ controller "gjarvis-bff/presentation/controller"
	usecase "gjarvis-bff/application/usecase"
)

func main() {
	log.Printf("Server started")
	r := chi.NewRouter()

	// todo add wire
	//sendMessageController := controller.NewSendMessageController(*dependency.InitSendMessageUsecase())
	sendMessageController.SetupSendMessageRoutes(r)

	log.Fatal(http.ListenAndServe("localhost:10000", r))
}