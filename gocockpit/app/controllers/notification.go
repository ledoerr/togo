package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"io"

	"github.com/ledoerr/togo/common/notification"
	"github.com/ledoerr/togo/gocockpit/app"
)

type Notification struct {
	*revel.Controller
}

func (c Notification) Heartbeat() revel.Result {

	heartbeat := decodeHeartbeat(c.Request.Body)

	app.UpdateServiceHeartbeat(heartbeat.Id, heartbeat.Time)

	return c.Render()
}

func decodeHeartbeat(c io.ReadCloser) notification.Heartbeat {
	dec := json.NewDecoder(c)

	heartbeat := notification.Heartbeat{}
	err := dec.Decode(&heartbeat)
	if err != nil {
		fmt.Errorf("heartbeat decoding failed: %v", err)
	}
	return heartbeat

}
