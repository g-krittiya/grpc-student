package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	pb "grpc-student/studentpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewStudentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetStudent(ctx, &pb.StudentRequest{
		Id: 200,
	})

	if err != nil {
		log.Fatalf("Error calling GetStudent: %v", err)
	}

	log.Printf("Student Info:")
	log.Printf("ID: %d", res.Id)
	log.Printf("Name: %s", res.Name)
	log.Printf("Major: %s", res.Major)
	log.Printf("Email: %s", res.Email)
	//log.Println("Phone: %s", res.Phone)

	students, err := client.ListStudents(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Error calling ListStudents: %v", err)
	}
	marshal, err := json.Marshal(students)
	if err != nil {
		log.Fatalf("Error marshalling students: %v", err)
		return
	}
	//log.Printf("Students: %v", students)
	log.Printf("Students: %s", marshal)
}
