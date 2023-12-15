package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/max-bazarov/go_chat_auth/pkg/auth_v1"
)

const grpcPort = 50051

type server struct {
	auth_v1.UnimplementedUserAPIV1Server
}

func (s *server) Create(ctx context.Context, req *auth_v1.CreateRequest) (*auth_v1.CreateResponse, error) {
	log.Printf("%#v", req.User)
	return &auth_v1.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Get(ctx context.Context, req *auth_v1.GetRequest) (*auth_v1.GetResponse, error) {
	log.Printf("Note id: %d", req.GetId())

	return &auth_v1.GetResponse{
		User: &auth_v1.User{
			Id: req.GetId(),
			Info: &auth_v1.UserInfo{
				Name:            gofakeit.Name(),
				Email:           gofakeit.Email(),
				Password:        gofakeit.BeerAlcohol(),
				PasswordConfirm: gofakeit.BeerAlcohol(),
				Role:            auth_v1.Role_USER,
			},
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

func (s *server) Update(ctx context.Context, req *auth_v1.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("%#v", req.Info)
	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *auth_v1.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	auth_v1.RegisterUserAPIV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
