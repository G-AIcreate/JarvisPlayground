package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	usecase "github.com/JarvisPlayground/gjarvis-bff/application/usecase"
	controller "github.com/JarvisPlayground/gjarvis-bff/presentation/controller"

	"github.com/go-chi/chi/v5"
)

func main() {
	log.Printf("Server started")
	
	// 環境変数の読み込み
	err := godotenv.Load(".env.test")
	if err != nil {
		log.Fatal("Error loading env file")
	}

	port := os.Getenv("PORT")
	log.Printf("BFF port %s", port)
	r := chi.NewRouter()

	sendMessageUsecase := usecase.NewSendMessageUsecase()
	sendMessageController := controller.NewSendMessageController(*sendMessageUsecase)
	sendMessageController.SetupSendMessageRoutes(r)

	http.ListenAndServe(port, r)
}
