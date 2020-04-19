package main

import (
	"context"
	"schedule-service/internal/pkg/api/schedule_service"
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

}
