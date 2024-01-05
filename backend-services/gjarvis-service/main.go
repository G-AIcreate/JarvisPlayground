// package main

// import (
// 	"context"
// 	"flag"
// 	"fmt"
// 	"log"
// 	"net"
// 	"time"

// 	pb "github.com/JarvisPlayground/gjarvis-service/proto"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/reflection"

// 	amqp "github.com/rabbitmq/amqp091-go"
// 	// "github.com/streadway/amqp"
// )

// var (
// 	port            = flag.Int("port", 50051, "The server port")
// 	messageChannel2 = make(chan string)
// 	messageChannel  = make(chan string)
// )

// // server is used to implement helloworld.GreeterServer.
// type server struct {
// 	pb.UnimplementedGJarvisServiceServer
// }

// // SayHello implements helloworld.GreeterServer
// func (s *server) JarvisHello(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
// 	log.Printf("Received: %v", req.GetName())
// 	return &pb.Reply{Message: "Hello " + req.GetName()}, nil
// }

// func clearChannel(ch chan string) {
// 	for {
// 		select {
// 		case <-ch:
// 			// 从通道中读取并丢弃消息
// 		default:
// 			// 通道为空时返回
// 			return
// 		}
// 	}
// }

// // gjarvisproto.GJarvisServiceServer の SendText メソッドの実現
// func (s *server) SendText(ctx context.Context, req *pb.TextRequest) (*pb.JarvisResponse, error) {
// 	log.Printf("backend/sendText: Received text message from bff: %s", req.GetTextMessage())

// 	messageChannel2 <- string(req.GetTextMessage())

// 	// テキストをカストマイズする処理
// 	// textAnswer := "Hello: " + req.GetTextMessage()
// 	// textAnswer := <-messageChannel
// 	// log.Printf("backend/sendText: get textAnswer from rabbitmq: %s", textAnswer)
// 	var textAnswer string
// 	// 使用select语句添加超时机制
// 	select {
// 	case textAnswer = <-messageChannel:
// 		log.Printf("backend/sendText: get textAnswer from rabbitmq: %s", textAnswer)
// 	case <-time.After(5 * time.Second): // 5秒超时
// 		log.Printf("backend/sendText: timeout waiting for response")
// 		textAnswer = "Default response due to timeout" // 超时后的默认响应
// 	}

// 	response := &pb.JarvisResponse{
// 		SessionId:   req.GetSessionId(),
// 		TextAnswer:  textAnswer,
// 		AudioAnswer: nil, // audio追加する
// 	}
// 	log.Printf("backend/sendText: return response to bff: %s", response)
// 	clearChannel(messageChannel)
// 	clearChannel(messageChannel2)
// 	return response, nil
// }

// // Error handling function
// func failOnError(err error, msg string) {
// 	if err != nil {
// 		log.Fatalf("%s: %s", msg, err)
// 		panic(fmt.Sprintf("%s: %s", msg, err))
// 	}
// }

// func main() {
// 	go Rabbit()

// 	flag.Parse()
// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
// 	if err != nil {
// 		log.Fatalf("backend/main: failed to listen: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	pb.RegisterGJarvisServiceServer(s, &server{})

// 	reflection.Register(s)

// 	log.Printf("server listening at %v", lis.Addr())
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("backend/main: failed to serve: %v", err)
// 	}

// }

// func Rabbit() {
// 	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
// 	failOnError(err, "Failed to connect to RabbitMQ")
// 	defer conn.Close()
// 	ch, err := conn.Channel()
// 	failOnError(err, "Failed to open a channel")
// 	defer ch.Close()

// 	// Queue for sending messages
// 	q, err := ch.QueueDeclare("go_to_python", false, false, false, false, nil)
// 	failOnError(err, "Failed to declare a queue")
// 	// Queue for receiving responses
// 	respQ, err := ch.QueueDeclare("python_to_go", false, false, false, false, nil)
// 	failOnError(err, "Failed to declare a response queue")

// 	// Publishing a message to 'go_to_python' queue
// 	textMessage := <-messageChannel2
// 	log.Printf("backend/rabbit: get textMessage from back-grpc-server: %s", textMessage)
// 	body := "Hello :" + textMessage

// 	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
// 		ContentType: "text/plain",
// 		Body:        []byte(body),
// 	})
// 	failOnError(err, "backend/rabbit: Failed to publish a message")

// 	// Consumer to receive response from Python
// 	msgs, err := ch.Consume(respQ.Name, "", true, false, false, false, nil)
// 	failOnError(err, "backend/rabbit: Failed to register a consumer")
// 	// forever := make(chan bool)
// 	go func() {
// 		for d := range msgs {
// 			log.Printf("backend/rabbit: Received a response from python: %s", d.Body)
// 			messageChannel <- string(d.Body)
// 		}
// 	}()
// 	log.Printf("backend/rabbit:  [*] Waiting for messages. To exit press CTRL+C")
// 	// <-forever

// }

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/google/uuid"

	pb "github.com/JarvisPlayground/gjarvis-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	port = flag.Int("port", 50051, "The server port")
	// RabbitMQ channels
	sendCh *amqp.Channel
	recvCh *amqp.Channel
	// Mutex for concurrent access to RabbitMQ
	mu            sync.Mutex
	correlationID string

	msgs <-chan amqp.Delivery
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
	textMessage := req.GetTextMessage()
	log.Printf("backend/sendText: Received text message from bff: %s", textMessage)

	// Send message to Python via RabbitMQ
	sendMessageToPython(textMessage)

	// Wait and receive response from Python
	responseFromPython := receiveMessageFromPython()

	response := &pb.JarvisResponse{
		SessionId:   req.GetSessionId(),
		TextAnswer:  responseFromPython,
		AudioAnswer: nil,
	}
	log.Printf("backend/sendText: return response to bff: %s", response)

	return response, nil
}

// Error handling function
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func generateUniqueID() string {
	// 生成一个唯一的ID，比如使用UUID
	return uuid.New().String()
}

func sendMessageToPython(message string) {
	mu.Lock()
	defer mu.Unlock()
	correlationID = generateUniqueID() // 生成唯一的 correlation ID
	log.Printf("send to python id: %s", correlationID)
	err := sendCh.Publish(
		"",             // exchange
		"go_to_python", // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			Body:          []byte(message),
			CorrelationId: correlationID,
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %s", err)
	}
}

func receiveMessageFromPython() string {
	log.Printf("call receive message from python method")

	mu.Lock()
	defer mu.Unlock()

	// 使用超时机制来等待新消息
	timeout := time.After(6 * time.Second)
	for {
		select {
		case d, ok := <-msgs:
			log.Printf("whatisthis！！！: %s", string(d.Body))
			if !ok {
				// 通道已关闭
				log.Println("RabbitMQ channel 'msgs' is closed")
				// closeRabbitMQResources(conn, sendCh, recvCh)
				// conn, sendCh, recvCh = initRabbitMQ()
				// msgs, err := recvCh.Consume(
				// 	"python_to_go", // queue
				// 	"",             // consumer
				// 	true,           // auto-ack
				// 	false,          // exclusive
				// 	false,          // no-local
				// 	false,          // no-wait
				// 	nil,            // args
				// )
				// failOnError(err, "Failed to re-register a consumer after reconnection")
				// continue
				return "Error: Channel closed"
			}

			// // 确认所有接收到的消息
			// if err := d.Ack(false); err != nil {
			// 	log.Printf("Error acknowledging message: %s", err)
			// }
			if d.CorrelationId == correlationID {
				//log.Println("sssssssssss: %s", string(d.Body))
				// 确认接收到的消息
				// if err := d.Ack(false); err != nil {
				// 	log.Printf("Error acknowledging message: %s", err)
				// }
				return string(d.Body)
			} else {
				log.Printf("elseelseelse: %s", string(d.Body))
			}
			//log.Println("xxxxxxxx: %s", string(d.Body))

		case <-timeout:
			log.Printf("Timeout waiting for response from Python")
			return "timeout"
		}
	}
}

func initRabbitMQ() (*amqp.Connection, *amqp.Channel, *amqp.Channel) {
	log.Printf("call init rabbit mq method")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	sendCh, err = conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	recvCh, err = conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}

	_, err = recvCh.QueueDeclare(
		"python_to_go", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// 初始化全局变量 msgs
	msgs, err = recvCh.Consume(
		"python_to_go", // queue
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	failOnError(err, "Failed to register a consumer")
	return conn, sendCh, recvCh
}

func closeRabbitMQResources(conn *amqp.Connection, sendCh, recvCh *amqp.Channel) {
	log.Printf("call close rabbit mq resources method")
	if err := sendCh.Close(); err != nil {
		log.Printf("Failed to close send channel: %s", err)
	}
	if err := recvCh.Close(); err != nil {
		log.Printf("Failed to close receive channel: %s", err)
	}
	if err := conn.Close(); err != nil {
		log.Printf("Failed to close RabbitMQ connection: %s", err)
	}
}

func main() {
	flag.Parse()
	conn, sendCh, recvCh := initRabbitMQ()

	// 在退出时清理资源
	defer closeRabbitMQResources(conn, sendCh, recvCh)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGJarvisServiceServer(s, &server{})
	reflection.Register(s)

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
