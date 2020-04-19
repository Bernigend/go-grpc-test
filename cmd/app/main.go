package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"schedule-service/internal/pkg/api/schedule_service"
)

const (
	ServerPort = 3456
)

type server struct {
	schedule_service.ScheduleServiceServer
}

func (s *server) Search(ctx context.Context, request *schedule_service.SearchRequest) (*schedule_service.SearchResponse, error) {
	// временная переменная
	var groups []*schedule_service.GroupItem

	// добавляем группу
	groups = append(groups, &schedule_service.GroupItem{
		Id:       "1",
		Name:     "Тест",
		Symbolic: "Test",
	})

	// возвращаем ответ
	return &schedule_service.SearchResponse{Groups: groups}, nil
}

func main() {
	flag.Parse()

	// начинаем прослушивать порт
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", ServerPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// create instance of grpc server
	grpcServer := grpc.NewServer()
	// register our service ]
	schedule_service.RegisterScheduleServiceServer(grpcServer, &server{})

	// запускаем обработку входящих запросов
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
