package setting

import "time"

type server struct {
	RunMode      string
	HttpPort     int64
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
