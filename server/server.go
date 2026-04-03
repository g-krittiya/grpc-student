package main

import (
	"context"
	"log"
	"net"

	pb "grpc-student/studentpb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStudentServiceServer
}

func (s *server) GetStudent(ctx context.Context, req *pb.StudentRequest) (*pb.StudentResponse, error) {

	log.Printf("Received request for student ID: %d", req.Id)

	// Mock data
	return &pb.StudentResponse{
		Id:    req.Id,
		Name:  "Alice Johnson",
		Major: "Computer Science",
		Email: "alice@university.com",
	}, nil
}

func (s *server) ListStudents(context.Context, *pb.Empty) (*pb.StudentListResponse, error) {
	log.Println("Received request for all students")

	listStudents := []*pb.StudentResponse{
		{
			Id:    100,
			Name:  "Mary Smith",
			Major: "Computer Science",
			Email: "marry.s@mail.com",
		},
		{
			Id:    101,
			Name:  "John Doe",
			Major: "Computer Science",
			Email: "john.d@mail.com",
		},
	}

	return &pb.StudentListResponse{Student: listStudents}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterStudentServiceServer(grpcServer, &server{})

	log.Println("gRPC Server running on port 8080")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
