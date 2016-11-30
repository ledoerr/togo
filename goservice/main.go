package main

import (
	"encoding/json"
	_ "expvar"
	"fmt"
	"github.com/ledoerr/togo/common/registration"
	"log"
	"net/http"
	"time"
	"bytes"
	"io"
	"os"
)

var (
	HEARTBEAT_PERIOD = time.Second * 5
	cockpitServer    = "http://localhost:9000"
	registerUrl      = "/registration/register"
	hostname         = "localhost"
	port             = "9123"
	name             = "foobar42"
	pollHandler      = "/debug/vars"
	pollUrl          = hostname + ":" + port + pollHandler
	pushUrl          = ""
)

func main() {

	go func() {
		err := http.ListenAndServe(hostname+":"+port, nil)
		if err != nil {
			fmt.Println(err)
		}
	}()

	registerAtServer()
	heartbeatPeriod := time.Tick(HEARTBEAT_PERIOD)
	for {
		select {
		case <-heartbeatPeriod:
			fmt.Println("foobar")
		}
	}
}

func registerAtServer() {

	request := generateRequest()
	response, err := http.Post(cockpitServer+registerUrl, "application/json", bytes.NewBuffer(request))
	defer response.Body.Close()
	if err != nil {
		log.Fatal("Request failed", err)
	}

	func(response *http.Response) {
		dec := json.NewDecoder(response.Body)
		responseMessage := common.RegisterResponse{}
		err = dec.Decode(&responseMessage)
		if err != nil {
			io.Copy(os.Stdout, response.Body)
			log.Fatal("Decoding failed:", err)
		}

		pushUrl = responseMessage.PushUrl
		fmt.Println("pushUrl:",pushUrl)
	}(response)

}
func generateRequest() []byte {
	request := common.RegisterRequest{Id: name, PollUrl: pollUrl}
	b, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	return b
}