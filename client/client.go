package main

import (
	"fmt"
	"io"
	"log"

	pb "example/Total_gRPC"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"bufio"
	"os"
	"strings"
)

const (
	address = "0.0.0.0:10023"
)

func queryUser(client pb.UserClient) {
	fmt.Println("client QueryUser")
	user, err := client.QueryUser(context.Background(), &pb.UserRequest{Uid: 1})
	if err != nil {
		log.Fatal("QueryInfo error: ", err)
	}
	fmt.Println("user: ", user)
	fmt.Printf("client QueryUser End\n\n")
}

func listUser(client pb.UserClient) {
	fmt.Println("client ListUser")
	stream, err := client.ListUser(context.Background(), &pb.UserCondition{
		Gender: 1,
	})
	if err != nil {
		log.Fatal("ListUser stream error: ", err)
	}
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("ListUser End\n\n")
			break
		}
		if err != nil {
			log.Fatal("ListUser error: ", err)
		}
		fmt.Println("user: ", user)
	}
}

func sendUser(client pb.UserClient) {
	fmt.Println("client sendUser")
	stream, err := client.SendUser(context.Background())
	if err != nil {
		log.Fatal("client SendUser create stream error: ", err)
	}
//1 mang cac con tro
	queries := []*pb.UserRequest{
		&pb.UserRequest{
			Uid: 1,
		},
		&pb.UserRequest{
			Uid: 2,
		},
		&pb.UserRequest{
			Uid: 3,
		},
		&pb.UserRequest{
			Uid: 4,
		},
		&pb.UserRequest{
			Uid: 5,
		},
	}
	for _, query := range queries {
		err := stream.Send(query)
		if err != nil {
			log.Fatal("SendUser error: ", err)
		}
	}
	summary, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("SendUser receive summary error: ", err)
	}
	fmt.Printf("Reveice summary: %s: %d\n", summary.GetDescription(), summary.GetTotal())
	fmt.Printf("SendUser End\n\n")
}

func chat(client pb.UserClient) {
	fmt.Println("Chat Start:")
	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatal("client Chat get stream error: ", err)
	}
	waitc := make(chan struct{})
	for  {
		talk, errr := stream.Recv()
		if errr == io.EOF {
			close(waitc)
			return
		}
		if errr != nil {
			log.Fatal("client Chat receive server chat error: ", err)
		}
		fmt.Println("server said: ", talk.GetTalk(), talk.GetDem())
		var a pb.UserMessage
		reader := bufio.NewReader(os.Stdin)
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)
		a.Talk = username

		err := stream.Send(&a)
		if err != nil {
			log.Fatal("client Chat send chat error: ", err)
		}
		talk.Dem++
	}
	err = stream.CloseSend()
	if err != nil {
		log.Fatal("client Chat close chat error: ", err)
	}
	<-waitc
	fmt.Printf("client Chat end\n\n")
}
func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("dail error:", err)
	}
	defer conn.Close()

	c := pb.NewUserClient(conn)

	queryUser(c)
	listUser(c)
	sendUser(c)
	chat(c)
}