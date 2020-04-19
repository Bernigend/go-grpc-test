package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"schedule-service/internal/pkg/api/schedule_service"
)

func main() {
	conn, err := grpc.Dial("localhost:3456")
	if err != nil {
		log.Fatalf("Error on grpc.Dial(): %v", err)
	}
	defer conn.Close()

	client := schedule_service.NewScheduleServiceClient(conn)
	group, err := client.Search(context.Background(), &schedule_service.SearchRequest{Name: "Test"})
	if err != nil {
		log.Fatalf("Failed to get group: %v", err)
	}

	log.Println(group)
}
