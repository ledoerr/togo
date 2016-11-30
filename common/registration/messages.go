package common

import (
	"time"
	"fmt"
)

type RegisterRequest struct {
	Id string `json:"id"`
	PollUrl string `json:"pollUrl"`
}

type RegisterResponse struct {
	Id string `json:"id"`
	PollUrl string `json:"pollUrl"`
	PushUrl string `json:"pushUrl"`
	Time time.Time `json:"time"`
}

func (r *RegisterRequest) String() string {
	return fmt.Sprintf("RegisterRequest: Id: %v PollUrl: %v", r.Id, r.PollUrl)
}


func (r *RegisterResponse) String() string{
	return fmt.Sprintf("RegisterResponse: Id: %v Time:%v PollUrl: %v PushUrl: %v", r.Id, r.Time, r.PollUrl, r.PushUrl)
}