package notification

import "time"

type Heartbeat struct {
	Id string `json:"id"`
	Time time.Time `json:"time"`
}

type Alert struct {
	Id string `json:"id"`
	InfoKey string `json:"infoKey"`
	InfoValue string `json:"infoValue"`
	Time time.Time `json:"time"`
}