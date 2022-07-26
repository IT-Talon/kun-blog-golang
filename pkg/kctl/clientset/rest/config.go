package rest

import "time"

type Config struct {
	Host    string
	Timeout time.Duration
}
