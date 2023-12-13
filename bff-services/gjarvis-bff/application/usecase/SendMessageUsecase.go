package usecase
import(
	// todo import problem needs to be fixed
	_ grpcclient "gjarvis-bff/application/proto"
	model "gjarvis-bff/application/model"
	pb "github.com/JarvisPlayground/gjarvis-bff/application/proto"
)
type SendMessageUsecase struct {
}

func SendMessageUsecase() *SendMessageUsecase {
	return &SendMessageUsecase
}

// SendTextToBackend 调用 gRPC 客户端
func (usecase *SendMessageUsecase) SendTextToBackend(textMessage string, sessionId string) (*model.JarvisResponse, error) {
    request := &pb.TextMessage {
		SessionId: sessionId,
		TextMessage: textMessage,
	}
	response, err := grpcclient.processTextMessage(request)
	if err != nil {
		return nil, err
	}
	answer := &model.JarvisResponse{
		TextAnswer: response.GetTextAnswer(),
		AudioAnswer: response.GetAudioAnswer(),
	}
	return answer;
}