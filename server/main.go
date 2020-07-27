package main

import (
    "context"
    "log"
    "net"
    "google.golang.org/grpc"
    pb "github.com/swiggy-test/grpc-demo/models"
    "strings"
)

type Server struct {
    pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, name *pb.Name) (*pb.GreeterResponse, error) {
    log.Printf("Recieved: %v", name)
    return &pb.GreeterResponse{Response: "Test response with the name " + name.GetFirstName() + " " + name.GetLastName()}, nil
}

func (s *Server) GreetFullName(ctx context.Context, name *pb.Name) (*pb.FullNameResponse, error) {
    log.Printf("Received: %v", name)
    return &pb.FullNameResponse{Name: &pb.Name{FirstName: strings.ToUpper(name.GetFirstName()), LastName: strings.ToUpper(name.GetLastName())}}, nil
}

func main() {
    log.Printf("Starting gRPC server\n")
    lis, err := net.Listen("tcp", "0.0.0.0:50000")
    if err != nil {
        log.Fatalf("Failed binding to interface: %v", err.Error())
    }
    s := grpc.NewServer()
    pb.RegisterGreeterServer(s, &Server{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to start gRPC server: %v", err.Error())
    }
}