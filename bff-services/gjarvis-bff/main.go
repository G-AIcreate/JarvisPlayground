package main

import (
	"log"
	"net/http"

	// todo import problem needs to be fixed
	controller "github.com/JarvisPlayground/gjarvis-bff/presentation/controller"

	usecase "github.com/JarvisPlayground/gjarvis-bff/application/usecase"

	"github.com/go-chi/chi/v5"
)

func main() {
	log.Printf("Server started")
	r := chi.NewRouter()

	sendMessageUsecase := usecase.NewSendMessageUsecase()
	sendMessageController := controller.NewSendMessageController(*sendMessageUsecase)
	sendMessageController.SetupSendMessageRoutes(r)

	log.Fatal(http.ListenAndServe("localhost:10000", r))
}
