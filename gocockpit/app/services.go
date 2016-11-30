package app

import (
	_ "fmt"
	"time"
	_ "sync"
)

type Service struct {
	Id 		string `json:"id"`
	Status	string `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	StatusUrl	string `json:"status_url"`
	Data	map[string]string `json:"data"`
}

var services = make(map[string]Service)

func init() {
	RegisterService("example.service");
}

func GetAllServices() []Service {

	list := make([]Service, 0, len(services))

	for _, service := range services {
		list = append(list, service)
	}

	return list
}

func RegisterService(id string) Service {

	service, exists := services[id]

	if(!exists) {
		service = Service{}
		service.Id = "luki"
		service.Status = "UP"
		service.StatusUrl = "http://some/endpoint"
		service.Timestamp = time.Now();
		services[id] = service
	}

	return service
}