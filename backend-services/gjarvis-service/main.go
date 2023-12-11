package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/JarvisPlayground/gjarvis-service/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGJarvisServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) JarvisHello(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.Reply{Message: "Hello " + in.GetName()}, nil
}

// gjarvisproto.GJarvisServiceServer の SendText メソッドの実現
func (s *server) SendText(ctx context.Context, in *pb.TextRequest) (*pb.JarvisResponse, error) {
	log.Printf("Received text message: %s", in.GetTextMessage())
	// テキストをカストマイズする処理
	textAnswer := "Processed: " + in.GetTextMessage()

	response := &pb.JarvisResponse{
		SessionId:   in.GetSessionId(),
		TextAnswer:  textAnswer,
		AudioAnswer: nil, // audio追加する
	}

	return response, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGJarvisServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
