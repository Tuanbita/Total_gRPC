package main

import (
	"fmt"
	"io"
	"log"
	"net"
	pb "example/Total_gRPC"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"

)

const (
	address = "0.0.0.0:10023"
)
//UserService struct la 1 su truu tuong cua server, no cho pheps dinh kem tai nguyen
//trong suot qua trinh call gRPC
type UserService struct {
	users []*pb.UserResponse
}
type TStringKey string
func (svr *UserService) QueryUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	fmt.Println("Query User")
	uid := req.GetUid()
	fmt.Println("User Id: ", uid)

	fmt.Printf("QueryUser end\n\n")

	return &pb.UserResponse{
		Name:   "Cloud",
		Age:    27,
		Gender: 0,
	}, nil
}

func (svr *UserService) ListUser(req *pb.UserCondition, stream pb.User_ListUserServer) error {
	fmt.Println("List User")
	gender := req.GetGender()
	fmt.Println("List User whose gender is: ", gender)

	for _, user := range svr.users {
		if user.Gender == gender {
			stream.Send(user)
		}
	}
	fmt.Printf("End List User\n\n")
	return nil
}
//sendUser: nhan request => stream user
func (svr *UserService) SendUser(stream pb.User_SendUserServer) error {
	fmt.Println("SendUser")

	var count uint32 = 0
	for {
		user, err := stream.Recv()
		fmt.Println("server receive user id ", user.GetUid())
		if err == io.EOF {
			fmt.Printf("server receive all users\n\n")
			//返回用户统计数据
			return stream.SendAndClose(&pb.UserSummary{
				Description: "total user",
				Total:       count,
			})
		}
		if err != nil {
			fmt.Println("server receive error: ", err)
		}
		count++
	}
}

func (svr *UserService) Chat(stream pb.User_ChatServer) error {
	fmt.Println("Begin Chat")
	for {
		stream.Send(&pb.UserMessage{
			Talk: "1-",
			Dem: 1,
		})
		in, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("chat done\n\n")
			return nil
		}

		if err != nil {
			fmt.Println("char error: ", err)
			return err
		}

		msg := in.GetTalk()
		dem := in.GetDem()
		fmt.Println("client chat: ", msg, dem)
		//回复
		stream.Send(&pb.UserMessage{
			Talk: "1234",
			Dem: dem ,
		})
	}
}
//---------------------------
//------------------------------
func main() {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	var User pb.UserResponse
/*
	Gender := testPoolRead("UserInfo", "Gender")
	Name := testPoolRead("UserInfo", "Name")
	Age := testPoolRead ("UserInfo","Age")

	age, err := strconv.ParseUint(Age, 10, 64)//mac dinh 64 bit hay sao day =))
	if err != nil {
		panic(err)
	}
	gender, err := strconv.ParseUint(Gender, 10, 64)
	if err != nil {
		panic(err)
	}

	User.Gender = uint32(gender)
	User.Name = Name
	User.Age = uint32(age)
*/
	svr := &UserService{
		users: []*pb.UserResponse{
			&User,
			&pb.UserResponse{
				Name:   "Jim",
				Age:    26,
				Gender: 0,
			},
			&pb.UserResponse{
				Name:   "Tom",
				Age:    10,
				Gender: 0,
			},
		},
	}

//tao 1 doi tuong server tren gRPC
	s := grpc.NewServer()
//dinh kem dich vu User vao trong server
	pb.RegisterUserServer(s, svr)

	fmt.Println("Listening on the 0.0.0.0:10023")

	//启动server
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}