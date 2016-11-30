package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/ledoerr/togo/common/registration"
	"github.com/revel/revel"
	"io"
	"time"
	"github.com/ledoerr/togo/gocockpit/app"
)

type Registration struct {
	*revel.Controller
}

const pushUrl = "notify"

func (c Registration) Register() revel.Result {

	registerRequest := decodeRequest(c.Request.Body)

	app.RegisterService(registerRequest.Id, registerRequest.PollUrl)

	registerResponse := common.RegisterResponse{Id: registerRequest.Id, Time: time.Now(), PollUrl: registerRequest.PollUrl, PushUrl: pushUrl}

	return c.RenderJson(registerResponse)
}

func decodeRequest(c io.ReadCloser) common.RegisterRequest {
	dec := json.NewDecoder(c)

	registerRequest := common.RegisterRequest{}
	err := dec.Decode(&registerRequest)
	if err != nil {
		fmt.Errorf("registerRequest decoding failed: %v", err)
	}
	return registerRequest

}
