package dtos

import "time"

type ChannelDTO struct {
	Name        string    `json:"name"`
	ReceiverUrl string    `json:"receiverUrl"`
	CreatedAt   time.Time `json:"createdAt"`
}

type CreateChannelDTO struct {
	Name string `json:"name"`
}
