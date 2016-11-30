package app

import (
	_ "fmt"
	"sync"
	"time"
)

type ServiceID string

type Service struct {
	Id        ServiceID         `json:"id"`
	Status    string            `json:"status"`
	Timestamp time.Time         `json:"timestamp"`
	StatusUrl string            `json:"status_url"`
	Data      map[string]string `json:"data"`
}

var services = make(map[string]Service)

var lock = sync.Mutex{}

func init() {
	RegisterService("example.service", "http://localhost:9999/health")
	RegisterService("sprong.service", "http://localhost:9100/health")
}

func GetAllServices() []Service {

	list := make([]Service, 0, len(services))

	lock.Lock()

	for _, service := range services {
		list = append(list, service)
	}

	lock.Unlock()

	return list
}

func RegisterService(id string, serviceUrl string) ServiceID {

	lock.Lock()

	service, exists := services[id]

	if !exists {
		service = Service{}
		service.Id = ServiceID(id)
		service.Status = "UNKNOWN"
		service.StatusUrl = serviceUrl
		service.Timestamp = time.Now()
		services[id] = service
	}

	lock.Unlock()

	return service.Id
}

func UpdateServiceStatus(id string, status string) Service {

	lock.Lock()

	service, exists := services[id]

	if exists {
		service.Status = status
		service.Timestamp = time.Now()
	}

	lock.Unlock()

	return service
}
