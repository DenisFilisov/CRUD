package model

import "time"

type KafkaMessage struct {
	Level    string
	Msg      string
	TimeSent time.Time
}
