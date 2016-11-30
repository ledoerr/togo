package main

import (
	"bytes"
	"encoding/json"
	_ "expvar"
	"fmt"
	"github.com/ledoerr/togo/common/registration"
	"github.com/ledoerr/togo/common/notification"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	HEARTBEAT_PERIOD = time.Second * 5
	cockpitServer    = "http://localhost:9000"
	registerUrl      = "/registration/register"
	hostname         = "localhost"
	port             = "9123"
	name             = "foobar42"
	pollHandler      = "/debug/vars"
	pollUrl          = "http://" + hostname + ":" + port + pollHandler
	pushUrl          = ""
	pushesHeartbeat  = true
)

func main() {

	go startHttpServer()
	registerAtCockpitServer()
	heartbeatPeriod := time.Tick(HEARTBEAT_PERIOD)
	for {
		select {
		case <-heartbeatPeriod:
			sendHeartbeat()
		}
	}
}
func startHttpServer() {
	err := http.ListenAndServe(hostname+":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func sendHeartbeat(){
	heartbeat := generateHeartbeat()
	_, err := http.Post(cockpitServer+"/"+pushUrl+"/heartbeat", "application/json", bytes.NewBuffer(heartbeat))
	if err != nil {
		log.Println("Sending heartbeat failed:", err)
	}
}

func generateHeartbeat() []byte {
	heartbeat := notification.Heartbeat{Id: name, Time:time.Now()}
	b, err := json.Marshal(heartbeat)
	if err != nil {
		log.Println(err)
	}
	return b
}

func registerAtCockpitServer() {

	request := generateRequest()
	response, err := http.Post(cockpitServer+registerUrl, "application/json", bytes.NewBuffer(request))
	defer response.Body.Close()
	if err != nil {
		log.Fatal("Request failed", err)
	}
	decodeResponse(response)
}

func generateRequest() []byte {
	request := registration.RegisterRequest{Id: name, PollUrl: pollUrl, PushesHeartbeat: pushesHeartbeat}
	b, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func decodeResponse(response *http.Response) {
	dec := json.NewDecoder(response.Body)
	responseMessage := registration.RegisterResponse{}
	err := dec.Decode(&responseMessage)
	if err != nil {
		io.Copy(os.Stdout, response.Body)
		log.Fatal("Decoding failed:", err)
	}

	pushUrl = responseMessage.PushUrl
	fmt.Println("pushUrl:", pushUrl)
}
