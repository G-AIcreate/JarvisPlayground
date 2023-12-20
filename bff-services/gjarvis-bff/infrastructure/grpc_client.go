package infrastructure

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/JarvisPlayground/gjarvis-bff/application/gjarvisproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	// client pb.GJarvisServiceClient
)

type GrpcClient struct{
	
}

func(i *GrpcClient) ProcessTextMessage(request *pb.TextRequest) (*pb.JarvisResponse, error) {
	// TODO switch to tls in PROD
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	// gRPC client作成
	client := pb.NewGJarvisServiceClient(conn)

	// request作る
	grpcRequest := &pb.TextRequest{TextMessage: request.TextMessage}

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// backend serviceを呼び出す
	response, err := client.SendText(ctx, grpcRequest)
	if err != nil {
		log.Fatalf("could not send text: %v", err)
		return nil, err
	}
	log.Printf("send text: %s", &response.TextAnswer)

	return response, nil
}
