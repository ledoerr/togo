package polling

import (
	"encoding/json"
	"github.com/ledoerr/togo/gocockpit/app"
	"net/http"
	"time"
	"fmt"
)

var polling_queue chan app.ServiceID = make(chan app.ServiceID)

func init() {
	go ping_scheduler(polling_queue)

	go pinger(polling_queue)
	go pinger(polling_queue)

}

func ping_scheduler(c chan app.ServiceID) {

	for {

		for _, service := range app.GetAllServices() {
			if service.IsPollable() {
				c <- service.Id
			}
		}

		time.Sleep(time.Second * 10)
	}
}

func pinger(c chan app.ServiceID) {
	for {
		// read from channel
		serviceId := <- c

		service, exists := app.GetServiceById(serviceId);
		if(exists) {
			fmt.Println("ping ", serviceId)

			status := ping_health(service.StatusUrl);

			app.UpdateServiceStatus(string(serviceId), status)
		}
	}
}


func ping_health(pollUrl string) string {
	var health Health

	resp, err := http.Get(pollUrl)
	if err != nil {
		health.Status = "DOWN (" + err.Error() + ")"
	} else {
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&health)
		if err != nil {
			health.Status = "DOWN (" + err.Error() + ")"
		}
	}

	return health.Status
}

type Health struct {
	Status string
}
