package usecase

import (
	// todo import problem needs to be fixed

	model "github.com/JarvisPlayground/gjarvis-bff/application/model"

	pb "github.com/JarvisPlayground/gjarvis-bff/application/gjarvisproto"

	infra "github.com/JarvisPlayground/gjarvis-bff/infrastructure"
)

type SendMessageUsecase struct {

}

func NewSendMessageUsecase() *SendMessageUsecase {
	return &SendMessageUsecase{}
}

// SendTextToBackend call gRPC client
func (usecase *SendMessageUsecase) SendTextToBackend(textMessage string, sessionId string) (*model.JarvisResponse, error) {
    request := &pb.TextRequest {
		SessionId: sessionId,
		TextMessage: textMessage,
	}
	i := infra.GrpcClient{}
	response, err := i.ProcessTextMessage(request)
	if err != nil {
		return nil, err
	}
	text := response.GetTextAnswer()
	audio := response.GetAudioAnswer()
	answer := &model.JarvisResponse{
		TextAnswer: model.TextAnswer(text),
		AudioAnswer: model.AudioAnswer(audio),
	}
	return answer, err;
}