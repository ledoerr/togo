package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/ledoerr/togo/common/registration"
	"github.com/ledoerr/togo/gocockpit/app"
	"github.com/revel/revel"
	"io"
	"time"
)

type Registration struct {
	*revel.Controller
}

const pushUrl = "notification"

func (c Registration) Register() revel.Result {

	registerRequest := decodeRequest(c.Request.Body)

	app.RegisterService(registerRequest.Id, registerRequest.PollUrl, registerRequest.PushesHeartbeat)

	registerResponse := registration.RegisterResponse{Id: registerRequest.Id, Time: time.Now(), PollUrl: registerRequest.PollUrl, PushUrl: pushUrl}

	return c.RenderJson(registerResponse)
}

func decodeRequest(c io.ReadCloser) registration.RegisterRequest {
	dec := json.NewDecoder(c)

	registerRequest := registration.RegisterRequest{}
	err := dec.Decode(&registerRequest)
	if err != nil {
		fmt.Errorf("registerRequest decoding failed: %v", err)
	}
	return registerRequest

}
