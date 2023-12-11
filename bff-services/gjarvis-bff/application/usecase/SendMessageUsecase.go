package usecase
import(
	// todo import problem needs to be fixed
	_ grpcclient "gjarvis-bff/application/proto"
)
type SendMessageUsecase struct {
}


// SendTextToBackend 调用 gRPC 客户端
func (usecase *SendMessageUsecase) SendTextToBackend(textMessage TextMessage) (ImplResponse, error) {
    return grpcclient.processTextMessage(textMessage)
}