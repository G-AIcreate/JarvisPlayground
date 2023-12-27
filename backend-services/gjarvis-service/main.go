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

	// amqp "github.com/rabbitmq/amqp091-go"
	"github.com/streadway/amqp"
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

var messageChannel2 = make(chan string) 

// gjarvisproto.GJarvisServiceServer の SendText メソッドの実現
func (s *server) SendText(ctx context.Context, req *pb.TextRequest) (*pb.JarvisResponse, error) {
	log.Printf("Received text message from bff: %s", req.GetTextMessage())

	messageChannel2 <- string(req.GetTextMessage())

	// テキストをカストマイズする処理
	// textAnswer := "Hello: " + req.GetTextMessage()
	textAnswer := <- messageChannel
	log.Printf("TextAnswer: %s", textAnswer)
	response := &pb.JarvisResponse{
		SessionId:   req.GetSessionId(),
		TextAnswer:  "test",
		AudioAnswer: nil, // audio追加する
	}
	log.Printf("Response: %s", response)
	return response, nil
}

// Error handling function
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	go Rabbit()

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

var messageChannel = make(chan string) 

func Rabbit() (){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Queue for sending messages
	q, err := ch.QueueDeclare("go_to_python", false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")
	// Queue for receiving responses
	respQ, err := ch.QueueDeclare("python_to_go", false, false, false, false, nil)
	failOnError(err, "Failed to declare a response queue")
	// Publishing a message to 'go_to_python' queue
	textMessage := <-messageChannel2
	log.Printf("TextMessage: %s", textMessage)
	body := "Hello :" + textMessage

		err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	failOnError(err, "Failed to publish a message")
	// Consumer to receive response from Python
	msgs, err := ch.Consume(respQ.Name, "", true, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a response from LLama2: %s", d.Body)
			messageChannel <- string(d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
