package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/JarvisPlayground/gjarvis-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGJarvisServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) JarvisHello(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	log.Printf("Received: %v", req.GetName())
	return &pb.Reply{Message: "Hello " + req.GetName()}, nil
}

// gjarvisproto.GJarvisServiceServer の SendText メソッドの実現
func (s *server) SendText(ctx context.Context, req *pb.TextRequest) (*pb.JarvisResponse, error) {
	log.Printf("Received text message: %s", req.GetTextMessage())
	// テキストをカストマイズする処理
	textAnswer := "Processed: " + req.GetTextMessage()

	response := &pb.JarvisResponse{
		SessionId:   req.GetSessionId(),
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

	reflection.Register(s);

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
