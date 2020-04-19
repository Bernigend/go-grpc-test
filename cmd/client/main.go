package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"schedule-service/internal/pkg/api/schedule_service"
)

func main() {
	// подключаемся к серверу
	conn, err := grpc.Dial("localhost:3456", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error on grpc.Dial(): %v", err)
	}
	defer conn.Close()

	// инициализируем клиент
	client := schedule_service.NewScheduleServiceClient(conn)

	// получаем ответ от апи
	group, err := client.Search(context.Background(), &schedule_service.SearchRequest{Name: "Test"})
	if err != nil {
		log.Fatalf("Failed to get group: %v", err)
	}

	log.Println(group)
}
