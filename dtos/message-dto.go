package dtos

import "time"

type ChannelMessageDTO struct {
	Channel    string              `json:"channel"`
	Headers    map[string][]string `json:"headers"`
	Body       []byte              `json:"body"`
	ReceivedAt time.Time           `json:"receivedAt"`
}
