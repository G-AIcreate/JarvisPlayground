package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	usecase "github.com/JarvisPlayground/gjarvis-bff/application/usecase"
	controller "github.com/JarvisPlayground/gjarvis-bff/presentation/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	log.Printf("Server started")

	// 環境変数の読み込み
	err := godotenv.Load(".env.test")
	if err != nil {
		log.Fatal("bff/main.go: Error loading env file")
	}

	port := os.Getenv("PORT")
	log.Printf("BFF port %s", port)
	r := chi.NewRouter()
	// CORS ミドルウェアの設定
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // すべてのオリジンを許可
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // 5分
	})
	// ルーターに CORS ミドルウェアを適用
	r.Use(cors.Handler)

	sendMessageUsecase := usecase.NewSendMessageUsecase()
	sendMessageController := controller.NewSendMessageController(*sendMessageUsecase)
	sendMessageController.SetupSendMessageRoutes(r)

	http.ListenAndServe(port, r)
}
