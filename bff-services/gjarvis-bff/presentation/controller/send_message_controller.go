package controller

import (
	"encoding/json"
	"net/http"

	// todo import problem needs to be fixed
	usecase "gjarvis-bff/application/usecase"

	_ "github.com/go-chi/chi/v5"
)

func (controller *SendMessageController) SetupSendMessageRoutes(r chi.Router) {
	// r.Get("/gjarvis/request_sessionId", controller.RequestSessionId)
	r.Post("/gjarvis/send_text", controller.SendText)
	// r.Post("/gjarvis/send_audio}", controller.SendAudio)
}

type SendMessageController struct {
	usecase usecase.SendMessageUsecase
}

func NewSendMessageController(usecase usecase.SendMessageUsecase) *SendMessageController {
	return &SendMessageController{usecase: usecase}
}

// SendText - Send text message to jarvis
func (s *SendMessageController) SendText(w http.ResponseWriter, r *http.Request) {
	// todo define TextMessage
	var textMessage TextMessage // 假设 TextMessage 已定义

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&textMessage); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// use usecase to call grpc lient
	response, err := s.usecase.SendTextToBackend(textMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// // GjarvisAPIService is a service that implements the logic for the GjarvisAPIServicer
// // This service should implement the business logic for every endpoint for the GjarvisAPI API.
// // Include any external packages or services that will be required by this service.
// type GjarvisAPIService struct {
// }

// // NewGjarvisAPIService creates a default api service
// func NewGjarvisAPIService() GjarvisAPIServicer {
// 	return &GjarvisAPIService{}
// }

// // RequestSessionId - Send sessionId request to jarvis
// func (s *GjarvisAPIService) RequestSessionId(ctx context.Context) (ImplResponse, error) {
// 	// TODO - update RequestSessionId with the required logic for this service method.
// 	// Add api_gjarvis_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

// 	// TODO: Uncomment the next line to return response Response(200, JarvisResponse{}) or use other options such as http.Ok ...
// 	// return Response(200, JarvisResponse{}), nil

// 	// TODO: Uncomment the next line to return response Response(405, {}) or use other options such as http.Ok ...
// 	// return Response(405, nil),nil

// 	return Response(http.StatusNotImplemented, nil), errors.New("RequestSessionId method not implemented")
// }

// // SendAudio - Send audio message to jarvis
// func (s *GjarvisAPIService) SendAudio(ctx context.Context, audioMessage AudioMessage) (ImplResponse, error) {
// 	// TODO - update SendAudio with the required logic for this service method.
// 	// Add api_gjarvis_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

// 	// TODO: Uncomment the next line to return response Response(200, JarvisResponse{}) or use other options such as http.Ok ...
// 	// return Response(200, JarvisResponse{}), nil

// 	return Response(http.StatusNotImplemented, nil), errors.New("SendAudio method not implemented")
// }
