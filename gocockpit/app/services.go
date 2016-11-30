package app

import (
	_ "fmt"
	"sort"
	"sync"
	"time"
)

type ServiceID string

type Service struct {
	Id               ServiceID         `json:"id"`
	Status           string            `json:"status"`
	HeartbeatTimeout time.Time         `json:"status"`
	Timestamp        time.Time         `json:"timestamp"`
	StatusUrl        string            `json:"status_url"`
	Data             map[string]string `json:"data"`
	PushesHeartbeat  bool              `json:"pushes_heartbeat"`
}

var services = make(map[ServiceID]Service)

var lock = sync.Mutex{}

func init() {
	go checkHeartbeat()
	RegisterService("example.service", "http://localhost:9999/health", false)
	RegisterService("sprong.service", "http://localhost:9100/health", false)
}

func checkHeartbeat() {
	for {
		allServices := GetAllServices()
		for _,service := range allServices {
			if service.PushesHeartbeat && !service.checkHeartbeat() {
				UpdateServiceStatus(string(service.Id), "DOWN: Heartbeat timeout")
			}
		}
	}
}

func (s *Service) checkHeartbeat() bool{
	return time.Now().Before(s.HeartbeatTimeout)
}


func GetAllServices() []Service {

	list := make([]Service, 0, len(services))
	var listKey []string

	lock.Lock()

	for _, service := range services {
		listKey = append(listKey, string(service.Id))

	}

	sort.Strings(listKey)
	for _, id := range listKey {
		list = append(list, services[ServiceID(id)])
	}

	lock.Unlock()

	return list
}

func RegisterService(id string, serviceUrl string, pushesHeartbeat bool) Service {

	lock.Lock()

	service, exists := services[ServiceID(id)]

	if !exists {
		serviceId := ServiceID(id)
		service = Service{}
		service.Id = serviceId
		service.Status = "UNKNOWN"
		service.StatusUrl = serviceUrl
		service.Timestamp = time.Now()
		service.PushesHeartbeat = pushesHeartbeat
		services[serviceId] = service
	}

	lock.Unlock()

	return service
}

func UpdateServiceStatus(id string, status string) Service {

	lock.Lock()

	service, exists := services[ServiceID(id)]

	if exists {
		service.Status = status
		service.Timestamp = time.Now()
	}

	lock.Unlock()

	return service
}

var heartbeatTimeout = time.Second * 30

func UpdateServiceHeartbeat(id string, heatbeatSentAt time.Time) {
	serviceId := ServiceID(id)

	lock.Lock()
	service, exists := services[serviceId]

	if exists {
		service.HeartbeatTimeout = heatbeatSentAt.Add(heartbeatTimeout)
		service.Timestamp = time.Now()
		service.Status = "UP"
		services[serviceId] = service
	}
	lock.Unlock()
}
